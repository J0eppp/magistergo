package magistergo

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"testing"
)

func TestMagisterGo(t *testing.T) {
	godotenv.Load(".env")
	accessToken := os.Getenv("ACCESSTOKEN")
	refreshToken := os.Getenv("REFRESHTOKEN")
	accessTokenExpires, _ := strconv.ParseInt(os.Getenv("EXPIRES"), 10, 64)
	tenant := os.Getenv("TENANT")
	magister, err := NewMagister(accessToken, refreshToken, accessTokenExpires, tenant)
	if err != nil {
		t.Failed()
		t.Error(err.Error())
	}

	res, err := magister.RefreshAccessToken()
	if err != nil {
		t.Failed()
		t.Error(err.Error())
	}

	os.Setenv("ACCESSTOKEN", res.AccessToken)
	os.Setenv("REFRESHTOKEN", res.RefreshToken)
	os.Setenv("EXPIRES", strconv.Itoa(int(res.ExpiresAt)))

	envMap, err := godotenv.Read(".env")
	envMap["ACCESSTOKEN"] = res.AccessToken
	envMap["REFRESHTOKEN"] = res.RefreshToken
	envMap["EXPIRES"] = strconv.Itoa(int(res.ExpiresAt))


	godotenv.Write(envMap, ".env")


	t.Logf("%+v\n", res)
}
