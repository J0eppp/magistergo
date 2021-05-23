package magistergo

import (
	"github.com/joho/godotenv"
	"os"
	"testing"
)

func TestMagisterGo(t *testing.T) {
	godotenv.Load()
	t.Log("Initializing Magister object")
	_, _ = NewMagister(os.Getenv("SCHOOL"), os.Getenv("USERNAME"), os.Getenv("PASSWORD"))
	//func() {
	//	magister.HTTPClient.Get("http://localhost:3000")
	//}()
}
