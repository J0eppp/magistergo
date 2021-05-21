package magistergo

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"magistergo/types"
	"net/http"
	"net/http/cookiejar"
	"strings"
	"time"
)

type Magister_I interface {

}

type Magister struct {
	School string
	Username string
	Password string
	Hostname string
	Authority string
	ClientID string
	RedirectURI string
	Scope string
	ResponseType string
	ACRValues string // Idk what this is
	DefaultState string
	DefaultNonce string
	AuthCode string
	Endpoints types.Endpoints
	HTTPClient http.Client
	//CookieJar *cookiejar.Jar
}

func NewMagister(school string, username string, password string) (Magister, error) {
	var magister Magister
	magister.School = school
	magister.Username = username
	magister.Password = password

	magister.Hostname = "https://" + strings.ToLower(magister.School) + ".magister.net/"

	magister.Authority = "https://accounts.magister.net"

	jar, err := cookiejar.New(&cookiejar.Options{})
	if err != nil {
		return magister, err
	}
	//magister.CookieJar = jar

	magister.HTTPClient = http.Client{
		Timeout: time.Second * 10,
		Jar: jar,
	}


	// Get the endpoints
	err = func() error {
		endpointsUrl := magister.Authority + "/.well-known/openid-configuration"
		res, err := http.Get(endpointsUrl)
		defer res.Body.Close()

		if err != nil {
			return err
		}

		endpointsBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}

		err = json.Unmarshal(endpointsBytes, &magister.Endpoints)
		if err != nil {
			return err
		}
		return nil
	}()
	if err != nil {
		return magister, err
	}

	magister.ClientID = "M6-" + magister.Hostname
	magister.RedirectURI = magister.Hostname + "/oid/redirect_callback.html"
	magister.Scope = "openid profile"
	magister.ResponseType = "id_token token"
	magister.ACRValues = "tenant:" + magister.Hostname
	magister.DefaultState = "00000000000000000000000000000000"
	magister.DefaultNonce = "00000000000000000000000000000000"


	// Get the cookies
	err = func() error {
		url := magister.Endpoints.AuthorizationEndpoint + "?client_id=" + magister.ClientID + "&redirect_uri=" + magister.RedirectURI + "&response_type=" + magister.ResponseType + "&scope=" + magister.Scope + "&acr_values=" + magister.ACRValues + "&state=" + magister.DefaultState + "&nonce=" + magister.DefaultNonce
		res, err := magister.HTTPClient.Get(url)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		//log.Println(res.Request.URL.String())

		resBytes, err := ioutil.ReadAll(res.Body)
		log.Println(string(resBytes))
		log.Printf("%+v\n", res)

		return nil
	}()


	// Get the authcode
	err = func() error {
		// Fetch the authcode
		authCodeUrl := "https://argo-web.vercel.app/api/authCode" // I'm fucked if this isn't supported anymore someday
		res, err := http.Get(authCodeUrl)
		defer res.Body.Close()
		if err != nil {
			return err
		}

		authCodeBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}

		magister.AuthCode = string(authCodeBytes)

		return nil
	}()
	if err != nil {
		return magister, err
	}

	return magister, nil
}