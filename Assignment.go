package magistergo

import (
	jsonitor "github.com/json-iterator/go"
	"io"
	"net/http"
	"strconv"
	"time"
)

// Assignment contains information about an assignment
type Assignment struct {
	ID    int `magisterjson:"Id" json:"id"`
	Links []struct {
		Rel  string `magisterjson:"Rel" json:"rel"`
		Href string `magisterjson:"Href" json:"href"`
	} `magisterjson:"Links"`
	Title                       string        `magisterjson:"Titel" json:"title"`
	Subject                     string        `magisterjson:"Vak" json:"subject"`
	HandInBefore                time.Time     `magisterjson:"InleverenVoor" json:"handInBefore"`
	HandedInAt                  interface{}   `magisterjson:"IngeleverdOp" json:"handedInAt"`
	StatusLastAssignmentVersion int           `magisterjson:"StatusLaatsteOpdrachtVersie" json:"statusLastAssignmentVersion"`
	LastAssignmentVersionNumber int           `magisterjson:"LaatsteOpdrachtVersienummer" json:"lastAssignmentVersionNumber"`
	Attachments                 []interface{} `magisterjson:"Bijlagen" json:"attachments"`
	Teachers                    interface{}   `magisterjson:"Docenten" json:"teachers"`
	VersionNavigationItems      []interface{} `magisterjson:"VersieNavigatieItems" json:"versionNavigationItems"`
	Description                 string        `magisterjson:"Omschrijving" json:"description"`
	Grading                     string        `magisterjson:"Beoordeling" json:"grading"`
	GradedAt                    time.Time     `magisterjson:"BeoordeeldOp" json:"gradedAt"`
	HandInAgain                 bool          `magisterjson:"OpnieuwInleveren" json:"handInAgain"`
	Closed                      bool          `magisterjson:"Afgesloten" json:"closed"`
	AllowedToHandIn             bool          `magisterjson:"MagInleveren" json:"allowedToHandIn"`
}

func (m *Magister) unmarshalAssignments(_assignments io.Reader) ([]Assignment, error) {
	var assignments []Assignment

	temp := struct {
		Items []Assignment `magisterjson:"Items"`
	}{}

	JSON := jsonitor.Config{
		TagKey: "magisterjson",
	}.Froze()

	err := JSON.NewDecoder(_assignments).Decode(&temp)
	assignments = temp.Items

	return assignments, err
}

// GetAssignments returns an array of assignments.
// The third parameter is there if you want to get more/less than 250 assignments.
func (m *Magister) GetAssignments(startDate time.Time, endDate time.Time, _amount ...int64) ([]Assignment, error) {
	var assignments []Assignment

	var amount int64
	if len(_amount) == 0 {
		amount = 250 // Get 250 assignments
	} else {
		amount = _amount[0]
	}

	if err := m.CheckSession(); err != nil {
		return assignments, err
	}

	begin := strconv.FormatInt(int64(startDate.Year()), 10) + "-" + strconv.FormatInt(int64(startDate.Month()), 10) + "-" + strconv.FormatInt(int64(startDate.Day()), 10)
	end := strconv.FormatInt(int64(endDate.Year()), 10) + "-" + strconv.FormatInt(int64(endDate.Month()), 10) + "-" + strconv.FormatInt(int64(endDate.Day()), 10)
	url := "https://" + m.Tenant + "/api/personen/" + m.UserID + "/opdrachten?skip=0&top=" + strconv.FormatInt(amount, 10) + "&einddatum=" + end + "&begindatum=" + begin + "&status=alle"

	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return assignments, err
	}

	r.Header.Add("authorization", "Bearer "+m.AccessToken)

	resp, err := m.HTTPClient.Do(r)
	if err != nil {
		return assignments, err
	}

	defer resp.Body.Close()

	assignments, err = m.unmarshalAssignments(resp.Body)
	if err != nil {
		return assignments, err
	}

	return assignments, nil
}
