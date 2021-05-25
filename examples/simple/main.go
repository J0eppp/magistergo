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

