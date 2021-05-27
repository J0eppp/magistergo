package main

import (
	"github.com/J0eppp/magistergo"
	"github.com/jasonlvhit/gocron"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

func main() {
	godotenv.Load("../../.env")
	expires, err := strconv.ParseInt(os.Getenv("EXPIRES"), 10, 64)
	if err != nil {
		panic(err)
	}
	
	magister, err := magistergo.NewMagister(os.Getenv("ACCESSTOKEN"), os.Getenv("REFRESHTOKEN"), expires, os.Getenv("TENANT"))
	if err != nil {
		panic(err)
	}

	err = gocron.Every(50).Minute().Do(magister.RefreshAccessToken())
	if err != nil {
		panic(err)
	}
}
