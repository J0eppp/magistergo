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

	// Create a Magister instance, give it all the data it needs
	magister, err := magistergo.NewMagister(accessToken, refreshToken, accessTokenExpires, tenant)
	if err != nil {
		fmt.Println(err)
		return
	}

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

	for _, message := range messages {
		fmt.Printf("%+v\n", message)
	}

	// Get the content of the last received message
	msgID := messages[0].ID
	message, err := magister.GetMessage(msgID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(message.Content)
}
```

## Examples
Find examples in the [examples folder](https://github.com/J0eppp/magistergo/tree/main/examples)