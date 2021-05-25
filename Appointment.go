package magistergo

import (
	"encoding/json"
	"net/http"
)

type Teacher struct {
	ID int64 `json:"Id"`
	Name string `json:"Naam"`
	TeacherCode string `json:"Docentcode"`
}

type Subject struct {
	ID int64 `json:"Id"`
	Name string `json:"Naam"`
}

type Classroom struct {
	Name string `json:"Naam"`
}

// TODO create an Attachment struct and add it to Appointment

// Appointment contains all the info about an appointment
type Appointment struct {
	ID int64 `json:"Id"`
	// Skip "Links", is that interesting to know?
	Start string `json:"Start"`
	End string `json:"Einde"`
	StartHour int8 `json:"LesuurVan"`
	EndHour int8 `json:"LesuurTotMet"`
	TakesWholeDay bool `json:"DuurtHeleDag"`
	Description string `json:"Omschrijving"`
	Location string `json:"Locatie"`
	Status int64 `json:"Status"`
	Type int64 `json:"Type"`
	IsOnline bool `json:"IsOnlineDeelname"`
	DisplayType int64 `json:"WeergaveType"`
	Content string `json:"Inhoud"`
	InfoType int64 `json:"InfoType"`
	Notes string `json:"Aantekeningen"`
	Finished bool `json:"Afgerond"`
	RepeatStatus int64 `json:"HerhaalStatus"`
	Repeat string `json:"Herhaling"`
	Subjects []Subject `json:"Vakken"`
	Teachers []Teacher `json:"Docenten"`
	Classrooms []Classroom `json:"Lokalen"`
	Groups string `json:"Groepen"`
	AssignmentID int64 `json:"OpdrachtId"`
	HasAttachements bool `json:"HeeftBijlagen"`
	Attachments string `json:"Bijlagen"`
}

func (appointment *Appointment) GetType() string {
	switch appointment.Type {
	case 0:   return "none" // None
	case 1:   return "personal" // Persoonlijk
	case 2:   return "general" // Algemeen
	case 3:   return "school wide" // School breed
	case 4:   return "internship" // Stage
	case 5:   return "intake" // Intake
	case 6:   return "free" // Roostervrij
	case 7:   return "kwt" // Kwt
	case 8:   return "standby" // Standby
	case 9:   return "blocked" // Blokkade
	case 10:  return "other" // Overig
	case 11:  return "blocked classroom" // Blokkade lokaal
	case 12:  return "blocked class" // Blokkade klas
	case 13:  return "class" // Les
	case 14:  return "study house" // Studiehuis
	case 15:  return "free study" // Roostervrije studie
	case 16:  return "schedule" // Planning
	case 101: return "measures" // Maatregelen
	case 102: return "presentations" // Presentaties
	case 103: return "exam schedule" // Examen rooster

	default: return "unknown"
	}
}

func (appointment *Appointment) GetInfoType() string {
	switch appointment.InfoType {
	case 0:  return "none" // None
	case 1:  return "homework" // Huiswerk
	case 2:  return "test" // Proefwerk
	case 3:  return "exam" // Tentamen
	case 4:  return "written exam" // Schriftelijke overhoring
	case 5:  return "oral exam" // Mondelinge overhoring
	case 6:  return "information" // Informatie
	case 7:  return "note" // Aantekening

	default: return "unknown"
	}
}

func (appointment *Appointment) GetStatus() string {
	switch appointment.Status {
	case 0:  return "unknown" // Geen status
	case 1:  return "scheduled automatically" // Geroosterd automatisch
	case 2:  return "scheduled manually" // Geroosterd handmatig
	case 3:  return "changed" // Gewijzigd
	case 4:  return "canceled manually" // Vervallen handmatig
	case 5:  return "canceled automatically" // Vervallen automatisch
	case 6:  return "in use" // In gebruik
	case 7:  return "finished" // Afgesloten
	case 8:  return "used" // Ingezet
	case 9:  return "moved" // Verplaatst
	case 10: return "changed and moved" // Gewijzigd en verplaatst

	default: return "unknown"
	}
}

func (magister *Magister) GetAppointments(dates... string) ([]Appointment, error) {
	var appointments []Appointment

	if err := magister.CheckSession(); err != nil {
		return appointments, err
	}

	var url string
	if len(dates) == 2 {
		url = "https://" + magister.Tenant + "/api/personen/" + magister.UserID + "/afspraken?status=1&tot=" + dates[1] + "&van=" + dates[0]
	} else {
		url = "https://" + magister.Tenant + "/api/personen/" + magister.UserID + "/afspraken"
	}


	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return appointments, err
	}

	r.Header.Add("authorization", "Bearer " + magister.AccessToken)

	resp, err := magister.HTTPClient.Do(r)
	if err != nil {
		return appointments, err
	}

	defer resp.Body.Close()

	temp := struct{
		Items []Appointment `json:"Items"`
	}{}

	err = json.NewDecoder(resp.Body).Decode(&temp)
	if err != nil {
		return appointments, err
	}

	appointments = temp.Items

	return appointments, nil
}