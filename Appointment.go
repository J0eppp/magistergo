package magistergo

import (
	jsonitor "github.com/json-iterator/go"
	"io"
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
	ID int64 `magisterjson:"Id" json:"id"`
	// Skip "Links", is that interesting to know?
	Start string `magisterjson:"Start" json:"start"`
	End string `magisterjson:"Einde" json:"end"`
	PeriodStarts int8 `magisterjson:"LesuurVan" json:"periodStarts"`
	PeriodEnds int8 `magisterjson:"LesuurTotMet" json:"periodEnds"`
	TakesWholeDay bool `magisterjson:"DuurtHeleDag" json:"TakesWholeDay"`
	Description string `magisterjson:"Omschrijving" json:"description"`
	Location string `magisterjson:"Locatie" json:"location"`
	Status int64 `magisterjson:"Status" json:"status"`
	Type int64 `magisterjson:"Type" json:"type"`
	IsOnline bool `magisterjson:"IsOnlineDeelname" json:"isOnline"`
	DisplayType int64 `magisterjson:"WeergaveType" json:"displayType"`
	Content string `magisterjson:"Inhoud" json:"content"`
	InfoType int64 `magisterjson:"InfoType" json:"infoType"`
	Notes string `magisterjson:"Aantekeningen" json:"notes"`
	Finished bool `magisterjson:"Afgerond" json:"finished"`
	RepeatStatus int64 `magisterjson:"HerhaalStatus" json:"repeatStatus"`
	Repeat string `magisterjson:"Herhaling" json:"repeat"`
	Subjects []Subject `magisterjson:"Vakken" json:"subjects"`
	Teachers []Teacher `magisterjson:"Docenten" json:"teachers"`
	Classrooms []Classroom `magisterjson:"Lokalen" json:"classrooms"`
	Groups string `magisterjson:"Groepen" json:"groups"`
	AssignmentID int64 `magisterjson:"OpdrachtId" json:"assignmentID"`
	HasAttachements bool `magisterjson:"HeeftBijlagen" json:"hasAttachments"`
	Attachments string `magisterjson:"Bijlagen" json:"attachments"`
}

// GetType returns the type of the appointment as a string
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

// GetInfoType returns the info type of the appointment as a string
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

// GetStatus returns the status of the appointment as a string
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

func (m *Magister) unmarshalAppointments(_appointments io.Reader) ([]Appointment, error) {
	var appointments []Appointment

	temp := struct {
		Items []Appointment `magisterjson:"Items"`
	}{}

	JSON := jsonitor.Config{
		TagKey: "magisterjson",
	}.Froze()

	err := JSON.NewDecoder(_appointments).Decode(&temp)
	appointments = temp.Items

	return appointments, err
}

// GetAppointments fetches the appointment data from Magister and puts it into a []Appointment
func (m *Magister) GetAppointments(dates... string) ([]Appointment, error) {
	var appointments []Appointment

	if err := m.CheckSession(); err != nil {
		return appointments, err
	}

	var url string
	if len(dates) == 2 {
		url = "https://" + m.Tenant + "/api/personen/" + m.UserID + "/afspraken?status=1&tot=" + dates[1] + "&van=" + dates[0]
	} else {
		url = "https://" + m.Tenant + "/api/personen/" + m.UserID + "/afspraken"
	}

	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return appointments, err
	}

	r.Header.Add("authorization", "Bearer " + m.AccessToken)

	resp, err := m.HTTPClient.Do(r)
	if err != nil {
		return appointments, err
	}

	defer resp.Body.Close()

	appointments, err = m.unmarshalAppointments(resp.Body)
	if err != nil {
		return appointments, err
	}

	return appointments, nil
}