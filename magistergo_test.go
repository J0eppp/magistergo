package magistergo

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestMagisterGo(t *testing.T) {
	// Get data from .env file
	err := godotenv.Load(".env")
	if err != nil {
		t.Failed()
		t.Error(err.Error())
	}

	accessToken := os.Getenv("ACCESSTOKEN")
	refreshToken := os.Getenv("REFRESHTOKEN")
	accessTokenExpires, _ := strconv.ParseInt(os.Getenv("EXPIRES"), 10, 64)
	tenant := os.Getenv("TENANT")

	t.Log("Testing on tenant: " + tenant)

	// // Environment files do not work on Windows
	// accessToken := ENV_ACCESS_TOKEN
	// refreshToken := ENV_REFRESH_TOKEN
	// accessTokenExpires := ENV_EXPIRES
	// tenant := ENV_TENANT

	// Create a Magister instance, give it all the data it needs
	t.Log("Creating a Magister instance")
	magister, err := NewMagister(accessToken, refreshToken, accessTokenExpires, tenant)
	if err != nil {
		t.Failed()
		t.Error(err.Error())
	}

	t.Log("Fetching appointments...")
	_, err = magister.GetAppointments()
	if err != nil {
		t.Failed()
		t.Error(err.Error())
	} else {
		t.Log("Fetched appointments!")
	}

	// Get appointments of one day
	t.Log("Fetching the appointments from a specific day...")
	_, err = magister.GetAppointments("2021-05-25", "2021-05-25")
	if err != nil {
		t.Failed()
		t.Error(err.Error())
	} else {
		t.Log("Fetched appointments!")
	}

	// Get messages
	t.Log("Fetching 7 last received messages...")
	messages, err := magister.GetMessages(7)
	if err != nil {
		t.Failed()
		t.Error(err)
	} else {
		t.Log("Fetched messages!")
	}

	// Get the content of the last received message
	t.Log("Fetching the content of the last received message...")
	msgID := messages[0].ID
	_, err = magister.GetMessage(msgID)
	if err != nil {
		t.Failed()
		t.Error(err)
	} else {
		t.Log("Fetched content of message!")
	}

	// Get the grades
	t.Log("Fetching grades...")
	_, err = magister.GetGrades()
	if err != nil {
		t.Failed()
		t.Error(err)
	} else {
		t.Log("Fetched grades!")
	}

	// Get the assignments
	t.Log("Fetching assignments...")
	_, err = magister.GetAssignments(time.Date(2020, 8, 1, 0, 0, 0, 0, &time.Location{}), time.Date(2021, 8, 1, 0, 0, 0, 0, &time.Location{}))
	if err != nil {
		t.Failed()
		t.Error(err)
	} else {
		t.Log("Fetched assignments!")
	}

	// Get absences
	t.Log("Fetching absences...")
	_, err = magister.GetAbsences()
	if err != nil {
		t.Failed()
		t.Error(err)
	} else {
		t.Log("Fetched absences!")
	}
}
