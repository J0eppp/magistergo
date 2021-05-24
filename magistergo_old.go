package magistergo
//
//import (
//	"crypto/tls"
//	"encoding/json"
//	client "github.com/bozd4g/go-http-client"
//	"io/ioutil"
//	"log"
//	"net/http"
//	"strings"
//	"time"
//)
//
//func NewMagister(school string, username string, password string) (Magister, error) {
//	var magister Magister
//	magister.School = school
//	magister.Username = username
//	magister.Password = password
//
//	magister.Tenant = strings.ToLower(magister.School) + ".magister.net"
//
//	magister.Authority = "https://accounts.magister.net"
//
//	magister.HTTPClient = http.Client{
//		Timeout: time.Second * 10,
//		//Jar: jar,
//		Transport: &http.Transport{
//			TLSNextProto: make(map[string]func(authority string, c *tls.Conn) http.RoundTripper),
//		},
//	}
//
//	// Get the endpoints
//	err := func() error {
//		endpointsUrl := magister.Authority + "/.well-known/openid-configuration"
//		res, err := http.Get(endpointsUrl)
//		defer res.Body.Close()
//
//		if err != nil {
//			return err
//		}
//
//		endpointsBytes, err := ioutil.ReadAll(res.Body)
//		if err != nil {
//			return err
//		}
//
//		err = json.Unmarshal(endpointsBytes, &magister.Endpoints)
//		if err != nil {
//			return err
//		}
//		return nil
//	}()
//	if err != nil {
//		return magister, err
//	}
//
//	magister.ClientID = "M6-" + magister.Tenant
//	magister.RedirectURI = "https://" + magister.Tenant + "/oid/redirect_callback.html"
//	magister.Scope = "openid profile"
//	magister.ResponseType = "id_token token"
//	magister.ACRValues = "tenant:" + magister.Tenant
//	magister.DefaultState = "00000000000000000000000000000000"
//	magister.DefaultNonce = "00000000000000000000000000000000"
//
//
//	// Get the cookies
//	err = func() error {
//		url := magister.Endpoints.AuthorizationEndpoint + "?client_id=M6LOAPP" /*+ magister.ClientID*/ + "&redirect_uri=" + magister.RedirectURI + "&response_type=" + magister.ResponseType + "&scope=" + magister.Scope + "&acr_values=" + magister.ACRValues + "&state=" + magister.DefaultState + "&nonce=" + magister.DefaultNonce
//		//url := "http://localhost:3000/connect/authorize" + "?client_id=" + magister.ClientID + "&redirect_uri=" + magister.RedirectURI + "&response_type=" + magister.ResponseType + "&scope=" + magister.Scope + "&acr_values=" + magister.ACRValues + "&state=" + magister.DefaultState + "&nonce=" + magister.DefaultNonce
//		log.Printf(url)
//		httpClient := client.New(url)
//		request, err := httpClient.Get("")
//		if err != nil {
//			return err
//		}
//
//		response, err := httpClient.Do(request)
//		if err != nil {
//			return err
//		}
//
//		log.Println(response.Get().Status)
//		log.Println(response.Get().Header)
//		log.Println(string(response.Get().Body))
//		//log.Printf("%+v\n", response)
//		//log.Println(url)
//		//res, err := magister.HTTPClient.Get(url)
//		//if err != nil {
//		//	return err
//		//}
//		//defer res.Body.Close()
//		////log.Println(res.Request.URL.String())
//		//
//		//resBytes, err := ioutil.ReadAll(res.Body)
//		//log.Println(string(resBytes))
//		//log.Printf("%+v\n", res)
//
//		return nil
//	}()
//
//
//	// Get the authcode
//	err = func() error {
//		// Fetch the authcode
//		authCodeUrl := "https://argo-web.vercel.app/api/authCode" // I'm fucked if this isn't supported anymore someday
//		res, err := http.Get(authCodeUrl)
//		defer res.Body.Close()
//		if err != nil {
//			return err
//		}
//
//		authCodeBytes, err := ioutil.ReadAll(res.Body)
//		if err != nil {
//			return err
//		}
//
//		magister.AuthCode = string(authCodeBytes)
//
//		return nil
//	}()
//	if err != nil {
//		return magister, err
//	}
//
//	return magister, nil
//}