package magistergo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func NewMagister(accessToken string, refreshToken string, accessTokenExpiresAt int64, tenant string) (Magister, error) {
	var magister Magister

	magister.Authority = "https://accounts.magister.net"
	magister.HTTPClient = http.Client{
		Timeout: time.Second * 10,
	}

		// Get the endpoints
		err := func() error {
			endpointsUrl := magister.Authority + "/.well-known/openid-configuration"
			res, err := http.Get(endpointsUrl)
			defer res.Body.Close()

			if err != nil {
				fmt.Errorf(err.Error())
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

	magister.AccessToken = accessToken
	magister.RefreshToken = refreshToken
	magister.AccessTokenExpiresAt = accessTokenExpiresAt
	magister.Tenant = tenant

	// Get AccountData
	err = func() error {
		if err := magister.CheckSession(); err != nil {
			return err
		}
		var user AccountData
		log.Println(magister.Tenant)
		url := "https://" + magister.Tenant + "/api/account?noCache=0"

		r, err := http.NewRequest(http.MethodGet, url, nil) // URL-encoded payload
		if err != nil {
			return  err
		}

		r.Header.Add("authorization", "Bearer " + magister.AccessToken)

		resp, err := magister.HTTPClient.Do(r)
		if err != nil {
			return err
		}

		defer resp.Body.Close()

		err = json.NewDecoder(resp.Body).Decode(&user)
		if err != nil {
			return err
		}

		magister.UserID = strconv.FormatInt(user.Persoon.Id, 10)

		return nil
	}()
	if err != nil {
		return magister, err
	}

	return magister, nil
}

// CheckSession checks if the session has expired
func (magister *Magister) CheckSession() error {
	//if time.Now().Unix() > magister.AccessTokenExpiresAt {
	//	return errors.New("your access token has expired")
	//}

	return nil
}

// RefreshAccessToken refreshes the access token
func (magister *Magister) RefreshAccessToken() (RefreshAccessTokenResponse, error) {
	var response RefreshAccessTokenResponse
	if err := magister.CheckSession(); err != nil {
		return response, err
	}

	data := url.Values{}
	data.Set("refresh_token", magister.RefreshToken)
	data.Set("client_id", "M6LOAPP")
	data.Set("grant_type", "refresh_token")

	r, err := http.NewRequest(http.MethodPost, magister.Endpoints.TokenEndpoint, strings.NewReader(data.Encode())) // URL-encoded payload
	if err != nil {
		return response, err
	}

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := magister.HTTPClient.Do(r)
	if err != nil {
		return response, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, err
	}

	response.ExpiresAt = time.Now().Unix() + response.expiresIn

	return response, nil
}