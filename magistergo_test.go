package magistergo

import (
	"github.com/joho/godotenv"
	"os"
	"testing"
)

func TestMagisterGo(t *testing.T) {
	godotenv.Load()
	magister, _ := NewMagister(os.Getenv("SCHOOL"), os.Getenv("USERNAME"), os.Getenv("PASSWORD"))

	t.Logf("%+v\n", magister.Endpoints)
}
