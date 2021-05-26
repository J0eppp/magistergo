package magistergo

import (
	jsonitor "github.com/json-iterator/go"
	"io"
	"time"
)

// TODO: finish this struct (the json name thingys)

type Grade struct {
	GradeID       int       `json:"gradeID" magisterjson:"CijferId"`
	GradeStr      string    `magisterjson:"CijferStr"`
	IsSufficient  bool      `magisterjson:"IsVoldoende"`
	EnteredBy     string    `magisterjson:"IngevoerdDoor"`
	DateEntered   time.Time `magisterjson:"DatumIngevoerd"`
	GradingPeriod struct {
		ID              int    `magisterjson:"Id"`
		Name            string `magisterjson:"Naam"`
		ReferenceNumber int    `magisterjson:"VolgNummer"`
	} `magisterjson:"CijferPeriode"`
	Vak struct {
		ID           int    `magisterjson:"Id"`
		Abbreviation string `magisterjson:"Afkorting"`
		Description  string `magisterjson:"Omschrijving"`
		FollowNumber int    `magisterjson:"Volgnr"`
	} `magisterjson:"Vak"`
	Retake     bool `magisterjson:"Inhalen"`
	Exemption  bool `magisterjson:"Vrijstelling"`
	Counts     bool `magisterjson:"TeltMee"`
	GadeColumn struct {
		ID                         int         `magisterjson:"Id"`
		ColumnName                 string      `magisterjson:"KolomNaam"`
		ColumnNumber               string      `magisterjson:"KolomNummer"`
		ColumnFollowNumber         string      `magisterjson:"KolomVolgNummer"`
		KolomKop                   string      `magisterjson:"KolomKop"`
		KolomOmschrijving          interface{} `magisterjson:"KolomOmschrijving"`
		KolomSoort                 int         `magisterjson:"KolomSoort"`
		IsHerkansingKolom          bool        `magisterjson:"IsHerkansingKolom"`
		IsDocentKolom              bool        `magisterjson:"IsDocentKolom"`
		HeeftOnderliggendeKolommen bool        `magisterjson:"HeeftOnderliggendeKolommen"`
		IsPTAKolom                 bool        `magisterjson:"IsPTAKolom"`
	} `magisterjson:"CijferKolom"`
	CijferKolomIdEloOpdracht int         `magisterjson:"CijferKolomIdEloOpdracht"`
	Docent                   interface{} `magisterjson:"Docent"`
	VakOntheffing            bool        `magisterjson:"VakOntheffing"`
	VakVrijstelling          bool        `magisterjson:"VakVrijstelling"`
}

// Todo: test this method
// https://stackoverflow.com/questions/38809137/golang-multiple-json-tag-names-for-one-field
func (m *Magister) unmarshalGrades(_grades io.Reader) ([]Grade, error) {
	var grades []Grade

	JSON := jsonitor.Config{
		TagKey: "magisterjson",
	}.Froze()

	err := JSON.NewDecoder(_grades).Decode(&grades)

	return grades, err
}