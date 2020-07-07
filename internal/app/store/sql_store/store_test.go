package sql_store_test

import (
	"fmt"
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M)  {
	databaseURL = os.Getenv("DATABASE_URL")

	fmt.Println(databaseURL)

	if databaseURL == "" {
		databaseURL = "user=root password=secret host=localhost dbname=restapi_test sslmode=disable"
	}

	os.Exit(m.Run())
}
