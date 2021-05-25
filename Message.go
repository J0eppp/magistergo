package magistergo

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type MessageInfo struct {
	ID        int    `json:"id"`
	Subject	  string `json:"onderwerp"`
	MapID     int    `json:"mapId"`
	Sender    struct {
		ID    int    `json:"id"`
		Name  string `json:"naam"`
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
		} `json:"links"`
	} `json:"afzender"`
	HasPriority		bool        `json:"heeftPrioriteit"`
	HasAttachments  bool        `json:"heeftBijlagen"`
	IsRead       	bool        `json:"isGelezen"`
	SentAt		    time.Time   `json:"verzondenOp"`
	ForwaredAt		interface{} `json:"doorgestuurdOp"`
	RepliedAt	    interface{} `json:"beantwoordOp"`
	Links           struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Map struct {
			Href string `json:"href"`
		} `json:"map"`
	} `json:"links"`
}

func (magister *Magister) GetMessages(amountOfMessages uint64, skip... uint64) ([]MessageInfo, error) {
	var messages []MessageInfo
	var skipMessages uint64

	if len(skip) == 0 {
		skipMessages = 0
	} else {
		skipMessages = skip[0]
	}

	url := "https://" + magister.Tenant + "/api/berichten/postvakin/berichten?top=" + strconv.FormatUint(amountOfMessages + skipMessages, 10) + "&skip=" + strconv.FormatUint(skipMessages, 10)

	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return messages, err
	}

	r.Header.Add("authorization", "Bearer " + magister.AccessToken)

	resp, err := magister.HTTPClient.Do(r)
	if err != nil {
		return messages, err
	}

	defer resp.Body.Close()

	temp := struct{
		Items []MessageInfo `json:"items"`
	}{}

	err = json.NewDecoder(resp.Body).Decode(&temp)
	if err != nil {
		return messages, err
	}

	messages = temp.Items

	return messages, nil
}

type Message struct {
	Content     string `json:"inhoud"`
	Receivers []struct {
		ID           int    `json:"id"`
		DisplayName	 string `json:"weergavenaam"`
		Type         string `json:"type"`
		IsToParent   bool   `json:"isAanOuder"`
		Links        struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
		} `json:"links"`
	} `json:"ontvangers"`
	CopyReceivers         []interface{} `json:"kopieOntvangers"`
	BlindCopyReceivers	  []interface{} `json:"blindeKopieOntvangers"`
	ID                    int           `json:"id"`
	Subject	              string        `json:"onderwerp"`
	MapID                 int           `json:"mapId"`
	Sender              struct {
		ID    int    `json:"id"`
		Name  string `json:"naam"`
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
		} `json:"links"`
	} `json:"afzender"`
	HasPriority		bool        `json:"heeftPrioriteit"`
	HasAttachments  bool        `json:"heeftBijlagen"`
	Isread	        bool        `json:"isGelezen"`
	SentAt		    time.Time   `json:"verzondenOp"`
	ForwardedAt	  	interface{} `json:"doorgestuurdOp"`
	RepliedAt	    interface{} `json:"beantwoordOp"`
	Links           struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Map struct {
			Href string `json:"href"`
		} `json:"map"`
		Attachments struct {
			Href string `json:"href"`
		} `json:"bijlagen"`
	} `json:"links"`
}