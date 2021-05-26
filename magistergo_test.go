package magistergo

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"testing"
)

func TestMagisterGo(t *testing.T) {
	// Get data from .env file
	err := godotenv.Load(".env")
	if err != nil {
		t.Failed()
		t.Error(err.Error())
	}

	accessToken := os.Getenv("ACCESS_TOKEN")
	refreshToken := os.Getenv("REFRESH_TOKEN")
	accessTokenExpires, _ := strconv.ParseInt(os.Getenv("EXPIRES"), 10, 64)
	tenant := os.Getenv("TENANT")

	// Create a Magister instance, give it all the data it needs
	magister, err := NewMagister(accessToken, refreshToken, accessTokenExpires, tenant)
	if err != nil {
		t.Failed()
		t.Error(err.Error())
	}

	_, err = magister.GetAppointments()
	if err != nil {
		t.Failed()
		t.Error(err.Error())
	}

	// Get appointments of one day
	_, err = magister.GetAppointments("2021-05-25", "2021-05-25")
	if err != nil {
		t.Failed()
		t.Error(err.Error())
	}

	// Get messages
	messages, err := magister.GetMessages(7)
	if err != nil {
		t.Failed()
		t.Error(err)
	}

	for _, message := range messages {
		t.Logf("%+v\n", message)
	}

	// Get the content of the last received message
	msgID := messages[0].ID
	message, err := magister.GetMessage(msgID)
	if err != nil {
		t.Failed()
		t.Error(err)
	}
	t.Log(message.Content)
}

