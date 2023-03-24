package kervan_test

import (
	"os"
	"testing"

	"github.com/kervandev/kervan-go"
)

func TestGetVersion(t *testing.T) {
	token := os.Getenv("KERVAN_TOKEN")
	api := kervan.NewCustomAPI("http://localhost:3010/api/v1", token)

	res, err := api.GetVersion()
	if err != nil {
		t.Error(err)
	}

	if res.Version == "" {
		t.Error("Version is empty")
	}

	if res.Description == "" {
		t.Error("Description is empty")
	}
}
