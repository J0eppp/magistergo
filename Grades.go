package magistergo

import (
	jsonitor "github.com/json-iterator/go"
	"io"
	"net/http"
	"strconv"
	"time"
)

// TODO: finish this struct (the json name thingys)

type Grade struct {
	GradeID       int       `magisterjson:"CijferId" json:"gradeID"`
	GradeStr      string    `magisterjson:"CijferStr" json:"gradeStr"`
	IsSufficient  bool      `magisterjson:"IsVoldoende" json:"isSufficient"`
	EnteredBy     string    `magisterjson:"IngevoerdDoor" json:"enteredBy"`
	DateEntered   time.Time `magisterjson:"DatumIngevoerd" json:"dateEntered"`
	GradingPeriod struct {
		ID              int    `magisterjson:"Id" json:"id"`
		Name            string `magisterjson:"Naam" json:"name"`
		ReferenceNumber int    `magisterjson:"VolgNummer" json:"referenceNumber"`
	} `magisterjson:"CijferPeriode" json:"gradingPeriod"`
	Subject struct {
		ID           int    `magisterjson:"Id" json:"id"`
		Abbreviation string `magisterjson:"Afkorting" json:"abbreviation"`
		Description  string `magisterjson:"Omschrijving" json:"description"`
		FollowNumber int    `magisterjson:"Volgnr" json:"followNumber"`
	} `magisterjson:"Vak" json:"subject"`
	CatchUp     bool `magisterjson:"Inhalen" json:"catchUp"`
	Exemption   bool `magisterjson:"Vrijstelling" json:"exemption"`
	Counts      bool `magisterjson:"TeltMee" json:"counts"`
	GradeColumn struct {
		ID                   int         `magisterjson:"Id" json:"id"`
		ColumnName           string      `magisterjson:"KolomNaam" json:"columnName"`
		ColumnNumber         string      `magisterjson:"KolomNummer" json:"columnNumber"`
		ColumnFollowNumber   string      `magisterjson:"KolomVolgNummer" json:"columnFollowNumber"`
		ColumnHeader         string      `magisterjson:"KolomKop" json:"columnHeader"`
		ColumnDescription    interface{} `magisterjson:"KolomOmschrijving" json:"columnDescription"`
		ColumnType           int         `magisterjson:"KolomSoort" json:"columnType"`
		IsRetakeColumn       bool        `magisterjson:"IsHerkansingKolom" json:"isRetakeColumn"`
		IsTeacherColumn      bool        `magisterjson:"IsDocentKolom" json:"isTeacherColumn"`
		HasUnderlyingColumns bool        `magisterjson:"HeeftOnderliggendeKolommen" json:"HasUnderlyingColumns"`
		IsPTAColumn          bool        `magisterjson:"IsPTAKolom" json:"isPTAColumn"`
	} `magisterjson:"CijferKolom" json:"gradeColumn"`
	GradeColumnIDEloAssignment int    `magisterjson:"CijferKolomIdEloOpdracht" json:"gradeColumnIDEloAssignment"`
	Teacher                    string `magisterjson:"Docent" json:"teacher"`
	SubjectExemption           bool   `magisterjson:"VakOntheffing" json:"subjectExemption"`
	SubjectDispensation        bool   `magisterjson:"VakVrijstelling" json:"subjectDispensation"`
}

type GradingPeriod struct {
	ID    int `magisterjson:"id" json:"id"`
	Study struct {
		ID    int    `magisterjson:"id" json:"id"`
		Code  string `magisterjson:"code" json:"code"`
		Links struct {
			Self struct {
				Href string `magisterjson:"href" json:"href"`
			} `magisterjson:"self" json:"self"`
		} `magisterjson:"links" json:"links"`
	} `magisterjson:"studie" json:"study"`
	Group struct {
		ID          int    `magisterjson:"id" json:"id"`
		Code        string `magisterjson:"code" json:"code"`
		Description string `magisterjson:"omschrijving" json:"description"`
		Links       struct {
			Self struct {
				Href string `magisterjson:"href" json:"href"`
			} `magisterjson:"self" json:"self"`
		} `magisterjson:"links" json:"links"`
	} `magisterjson:"groep" json:"group"`
	ClassPeriod struct {
		Code  string `magisterjson:"code" json:"code"`
		Links struct {
			Self struct {
				Href string `magisterjson:"href" json:"href"`
			} `magisterjson:"self" json:"self"`
		} `magisterjson:"links" json:"links"`
	} `magisterjson:"lesperiode" json:"classPeriod"`
	StudyProgram []struct {
		Code  string `magisterjson:"code" json:"code"`
		Links struct {
			Self struct {
				Href string `magisterjson:"href" json:"href"`
			} `magisterjson:"self" json:"self"`
		} `magisterjson:"links" json:"links"`
	} `magisterjson:"profielen" json:"studyProgram"`
	PersonalMentor struct {
		Initials      string `magisterjson:"voorletters" json:"initials"`
		SurnamePrefix string `magisterjson:"tussenvoegsel" json:"surnamePrefix"`
		Lastname      string `magisterjson:"achternaam" json:"lastName"`
		Links         struct {
			Self struct {
				Href string `magisterjson:"href" json:"href"`
			} `magisterjson:"self" json:"self"`
		} `magisterjson:"links" json:"links"`
	} `magisterjson:"persoonlijkeMentor" json:"personalMentor"`
	Start              string `magisterjson:"begin" json:"start"`
	End                string `magisterjson:"einde" json:"end"`
	IsMainRegistration bool   `magisterjson:"isHoofdAanmelding" json:"isMainRegistration"`
	Links              struct {
		Self struct {
			Href string `magisterjson:"href" json:"href"`
		} `magisterjson:"self" json:"self"`
		Subjects struct {
			Href string `magisterjson:"href" json:"href"`
		} `magisterjson:"vakken" json:"subjects"`
		Periods struct {
			Href string `magisterjson:"href" json:"href"`
		} `magisterjson:"perioden" json:"periods"`
		Grades struct {
			Href string `magisterjson:"href" json:"href"`
		} `magisterjson:"cijfers" json:"grades"`
		Mentors struct {
			Href string `magisterjson:"href" json:"href"`
		} `magisterjson:"mentoren" json:"mentors"`
	} `magisterjson:"links" json:"links"`
}

func (m *Magister) unmarshalGradingPeriods(_gradingPeriods io.Reader) ([]GradingPeriod, error) {
	var gradingPeriods []GradingPeriod

	temp := struct {
		Items []GradingPeriod `magisterjson:"items"`
	}{}

	JSON := jsonitor.Config{
		TagKey: "magisterjson",
	}.Froze()

	err := JSON.NewDecoder(_gradingPeriods).Decode(&temp)
	gradingPeriods = temp.Items

	return gradingPeriods, err
}

func (m *Magister) GetGrades() ([]Grade, error) {
	var grades []Grade

	if err := m.CheckSession(); err != nil {
		return grades, err
	}

	// Get the grading periods (years)
	begin := strconv.FormatInt(int64(time.Now().AddDate(-10, 0, 0).Year()), 10) + "-" + strconv.FormatInt(int64(time.Now().AddDate(-10, 0, 0).Month()), 10) + "-" + strconv.FormatInt(int64(time.Now().AddDate(-10, 0, 0).Day()), 10)
	end := strconv.FormatInt(int64(time.Now().AddDate(1, 6, 0).Year()), 10) + "-" + strconv.FormatInt(int64(time.Now().AddDate(1, 6, 0).Month()), 10) + "-" + strconv.FormatInt(int64(time.Now().AddDate(1, 6, 0).Day()), 10)
	url := "https://" + m.Tenant + "/api/leerlingen/" + m.UserID + "/aanmeldingen?begin=" + begin + "&einde=" + end
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return grades, err
	}

	r.Header.Add("authorization", "Bearer "+m.AccessToken)

	resp, err := m.HTTPClient.Do(r)
	if err != nil {
		return grades, err
	}

	defer resp.Body.Close()

	gradingPeriods, err := m.unmarshalGradingPeriods(resp.Body)
	lastGradingPeriod := gradingPeriods[0]

	url = "https://" + m.Tenant + "/api/personen/" + m.UserID + "/aanmeldingen/" + strconv.FormatInt(int64(lastGradingPeriod.ID), 10) + "/cijfers/cijferoverzichtvooraanmelding?actievePerioden=true&alleenBerekendeKolommen=false&alleenPTAKolommen=false"

	r, err = http.NewRequest("GET", url, nil)
	if err != nil {
		return grades, err
	}

	r.Header.Add("authorization", "Bearer "+m.AccessToken)

	resp, err = m.HTTPClient.Do(r)
	if err != nil {
		return grades, err
	}

	defer resp.Body.Close()

	_grades, err := m.unmarshalGrades(resp.Body)
	if err != nil {
		return grades, err
	}
	grades = _grades

	return grades, nil
}

// https://stackoverflow.com/questions/38809137/golang-multiple-json-tag-names-for-one-field
func (m *Magister) unmarshalGrades(_grades io.Reader) ([]Grade, error) {
	var grades []Grade

	temp := struct {
		Items []Grade `magisterjson:"Items"`
	}{}

	JSON := jsonitor.Config{
		TagKey: "magisterjson",
	}.Froze()

	err := JSON.NewDecoder(_grades).Decode(&temp)
	grades = temp.Items

	return grades, err
}
