package magistergo

import (
	jsonitor "github.com/json-iterator/go"
	"io"
	"net/http"
	"strconv"
	"time"
)

type Absence struct {
	ID                 int         `magisterjson:"Id" json:"id"`
	Start              time.Time   `magisterjson:"Start" json:"start"`
	End                time.Time   `magisterjson:"Eind" json:"end"`
	Period             int         `magisterjson:"Lesuur" json:"period"`
	Allowed            bool        `magisterjson:"Geoorloofd" json:"allowed"`
	AppointmentID      int         `magisterjson:"AfspraakId" json:"appointmentID"`
	Description        string      `magisterjson:"Omschrijving" json:"description"`
	AccountabilityType int         `magisterjson:"Verantwoordingtype" json:"accountabilityType"`
	Code               string      `magisterjson:"Code" json:"code"`
	Appointment        Appointment `magisterjson:"Afspraak" json:"appointment"`
	// Afspraak           struct {
	// 	Id               int         `magisterjson:"Id"`
	// 	Links            interface{} `magisterjson:"Links"`
	// 	Start            time.Time   `magisterjson:"Start"`
	// 	Einde            time.Time   `magisterjson:"Einde"`
	// 	LesuurVan        int         `magisterjson:"LesuurVan"`
	// 	LesuurTotMet     int         `magisterjson:"LesuurTotMet"`
	// 	DuurtHeleDag     bool        `magisterjson:"DuurtHeleDag"`
	// 	Omschrijving     string      `magisterjson:"Omschrijving"`
	// 	Lokatie          string      `magisterjson:"Lokatie"`
	// 	Status           int         `magisterjson:"Status"`
	// 	Type             int         `magisterjson:"Type"`
	// 	IsOnlineDeelname bool        `magisterjson:"IsOnlineDeelname"`
	// 	WeergaveType     int         `magisterjson:"WeergaveType"`
	// 	Inhoud           string      `magisterjson:"Inhoud"`
	// 	InfoType         int         `magisterjson:"InfoType"`
	// 	Aantekening      interface{} `magisterjson:"Aantekening"`
	// 	Afgerond         bool        `magisterjson:"Afgerond"`
	// 	HerhaalStatus    int         `magisterjson:"HerhaalStatus"`
	// 	Herhaling        interface{} `magisterjson:"Herhaling"`
	// 	Vakken           []struct {
	// 		Id   int         `magisterjson:"Id"`
	// 		Naam interface{} `magisterjson:"Naam"`
	// 	} `magisterjson:"Vakken"`
	// 	Docenten interface{} `magisterjson:"Docenten"`
	// 	Lokalen  []struct {
	// 		Naam interface{} `magisterjson:"Naam"`
	// 	} `magisterjson:"Lokalen"`
	// 	Groepen       interface{} `magisterjson:"Groepen"`
	// 	OpdrachtId    int         `magisterjson:"OpdrachtId"`
	// 	HeeftBijlagen bool        `magisterjson:"HeeftBijlagen"`
	// 	Bijlagen      interface{} `magisterjson:"Bijlagen"`
	// } `magisterjson:"Afspraak"`
}

type AbsencePeriod struct {
	Start       time.Time `magisterjson:"Start" json:"start"`
	End         time.Time `magisterjson:"Eind" json:"end"`
	Description string    `magisterjson:"Omschrijving" json:"description"`
}

func (m *Magister) unmarshalAbsencePeriods(_absencePeriods io.Reader) ([]AbsencePeriod, error) {
	var absencePeriods []AbsencePeriod

	temp := struct {
		Items []AbsencePeriod `magisterjson:"Items"`
	}{}

	JSON := jsonitor.Config{
		TagKey: "magisterjson",
	}.Froze()

	err := JSON.NewDecoder(_absencePeriods).Decode(&temp)
	absencePeriods = temp.Items

	return absencePeriods, err
}

func (m *Magister) unmarshalAbsences(_absences io.Reader) ([]Absence, error) {
	var absences []Absence

	if err := m.CheckSession(); err != nil {
		return absences, err
	}

	temp := struct {
		Items []Absence `magisterjson:"Items"`
	}{}

	JSON := jsonitor.Config{
		TagKey: "magisterjson",
	}.Froze()

	err := JSON.NewDecoder(_absences).Decode(&temp)
	absences = temp.Items

	return absences, err
}

func (m *Magister) GetAbsences(dates... time.Time) ([]Absence, error) {
	var absences []Absence
	
	var _begin time.Time
	var _end time.Time

	if len(dates) > 2 {
		_begin = dates[0]
		_end = dates[1]
	} else {
		var absencePeriods []AbsencePeriod

		url := "https://" + m.Tenant + "/api/personen/" + m.UserID + "/absentieperioden"

		r, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return absences, err
		}

		r.Header.Add("authorization", "Bearer "+m.AccessToken)

		resp, err := m.HTTPClient.Do(r)
		if err != nil {
			return absences, err
		}

		defer resp.Body.Close()

		absencePeriods, err = m.unmarshalAbsencePeriods(resp.Body)
		if err != nil {
			return absences, err
		}

		_begin = absencePeriods[0].Start
		_end = absencePeriods[0].End
	}

	begin := strconv.FormatInt(int64(_begin.Year()), 10) + "-" + strconv.FormatInt(int64(_begin.Month()), 10) + "-" + strconv.FormatInt(int64(_begin.Day()), 10)
	end := strconv.FormatInt(int64(_end.Year()), 10) + "-" + strconv.FormatInt(int64(_end.Month()), 10) + "-" + strconv.FormatInt(int64(_end.Day()), 10)


	url := "https://" + m.Tenant + "/api/personen/" + m.UserID + "/absenties?tot=" + end + "&van=" + begin

	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return absences, err
	}

	r.Header.Add("authorization", "Bearer "+m.AccessToken)

	resp, err := m.HTTPClient.Do(r)
	if err != nil {
		return absences, err
	}

	defer resp.Body.Close()

	absences, err = m.unmarshalAbsences(resp.Body)

	return absences, err
}
