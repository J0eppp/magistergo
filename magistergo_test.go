package magistergo

import (
	"testing"
)

func TestMagisterGo(t *testing.T) {
	//// Get data from .env file
	//err := godotenv.Load()
	//if err != nil {
	//	t.Failed()
	//	t.Error(err.Error())
	//}


	//accessToken := os.Getenv("ACCESS_TOKEN")
	//refreshToken := os.Getenv("REFRESH_TOKEN")
	//accessTokenExpires, _ := strconv.ParseInt(os.Getenv("EXPIRES"), 10, 64)
	//tenant := os.Getenv("TENANT")


	// Environment files do not work on Windows
	accessToken := ENV_ACCESS_TOKEN
	refreshToken := ENV_REFRESH_TOKEN
	accessTokenExpires := ENV_EXPIRES
	tenant := ENV_TENANT

	// Create a Magister instance, give it all the data it needs
	magister, err := NewMagister(accessToken, refreshToken, accessTokenExpires, tenant)
	if err != nil {
		t.Failed()
		t.Error(err.Error())
	}

	appointments, err := magister.GetAppointments()
	if err != nil {
		t.Failed()
		t.Error(err.Error())
	}
	t.Log(appointments[0].Description)

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

	// Get the content of the last received message
	msgID := messages[0].ID
	_, err = magister.GetMessage(msgID)
	if err != nil {
		t.Failed()
		t.Error(err)
	}

	// Get the grades
	_, err = magister.GetGrades()
	if err != nil {
		t.Failed()
		t.Error(err)
	}
}

