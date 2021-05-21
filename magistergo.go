package magistergo

import (
	"encoding/json"
	"io/ioutil"
	"magistergo/types"
	"net/http"
)

type Magister_I interface {

}

type Magister struct {
	AuthCode string
	Endpoints types.Endpoints
	HTTPClient http.Client
	CookieJar *http.CookieJar
}

func NewMagister(school string, username string, password string) (Magister, error) {
	var magister Magister

	authority := "https://accounts.magister.net"

	// Get the endpoints
	err := func() error {
		endpointsUrl := authority + "/.well-known/openid-configuration"
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





	// Get the authcode
	err = func() error {
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