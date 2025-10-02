package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "user=andrew host=localhost dbname=restapi_test sslmode=disable password=102309c"
	}

	os.Exit(m.Run())
}
