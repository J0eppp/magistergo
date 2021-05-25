# magistergo
A go implementation of the Magister 6 API

# Usage
## Get the tokens using an external library
Get the session token and refresh token using another library (e.g. [magister-scraper](https://github.com/JipFr/magister-scraper/))
```javascript
const { AuthManager } = require("magister-openid");
const axios = require("axios");

const options = {
    tenant: "school.magister.net",
    username: "<username>",
    password: "<password>"
}

const manager = new AuthManager(options.tenant);

axios("https://argo-web.vercel.app/api/authCode").then(({ data: authCode }) => {
    manager.login(options.username, options.password, authCode).then(tokens => {
        console.log("ACCESSTOKEN=" + tokens.access_token);
        console.log("REFRESHTOKEN=" + tokens.refresh_token);
        console.log("EXPIRES=" + tokens.expires_at);
        console.log("TENANT=" + options.tenant);
    });
});
```

<b>*Note: this library is not able to get the access and refresh token by itself (yet), however it is able to refresh the access token with the refresh token.*</b> 

## Use the library
```go
package main

import (
	"fmt"
	"github.com/J0eppp/magistergo"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

func main() {
	// Get data from .env file
	godotenv.Load(".env")
	accessToken := os.Getenv("ACCESSTOKEN")
	refreshToken := os.Getenv("REFRESHTOKEN")
	accessTokenExpires, _ := strconv.ParseInt(os.Getenv("EXPIRES"), 10, 64)
	tenant := os.Getenv("TENANT")
	magister, err := magistergo.NewMagister(accessToken, refreshToken, accessTokenExpires, tenant)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Refresh the access token
	res, err := magister.RefreshAccessToken()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Safe the new access and refresh token to the .env file
	os.Setenv("ACCESSTOKEN", res.AccessToken)
	os.Setenv("REFRESHTOKEN", res.RefreshToken)
	os.Setenv("EXPIRES", strconv.Itoa(int(res.ExpiresAt)))

	envMap, err := godotenv.Read(".env")
	envMap["ACCESSTOKEN"] = res.AccessToken
	envMap["REFRESHTOKEN"] = res.RefreshToken
	envMap["EXPIRES"] = strconv.Itoa(int(res.ExpiresAt))
	envMap["TENANT"] = tenant


	godotenv.Write(envMap, ".env")

	// Get appointments
	_, err = magister.GetAppointments()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get today's appointments
	_, err = magister.GetAppointments("2021-05-25", "2021-05-25")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get messages
	messages, err := magister.GetMessages(7)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print 7 messages
	for _, message := range messages {
		fmt.Printf("%+v\n", message)
	}
}
```

## Examples
Find examples in the [examples folder](https://github.com/J0eppp/magistergo/tree/main/examples)