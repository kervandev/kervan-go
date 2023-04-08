package kervan_test

import (
	"os"
	"testing"

	"github.com/kervandev/kervan-go"
)

var (
	testToken     = os.Getenv("TOKEN")
	productSecret = os.Getenv("PRODUCT_SECRET")
)

func TestGetVersion(t *testing.T) {
	token := testToken
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

func TestCheckLicence(t *testing.T) {
	token := testToken
	api := kervan.NewCustomAPI("http://localhost:3010/api/v1", token)

	gtoken, err := kervan.GenerateLicenceCheckJWT("ed9aba8e-cd2e-4af8-b265-1494997a05a8", "192.168.1.1", map[string]string{"test": "test"}, productSecret)
	if err != nil {
		t.Error(err)
	}

	payload := kervan.CheckLicencePayload{
		Token: gtoken,
	}

	res, err := api.CheckLicence(&payload)
	if err != nil {
		t.Error(err)
	}

	claims, err := kervan.ParseLicenceCheckResponseJWT(res.Token, productSecret)
	if err != nil {
		t.Error(err)
	}

	if claims.IsValid == false {
		t.Error("IsValid is false")
	}

	if claims.PlanCode == "" {
		t.Error("PlanCode is empty")
	}

}
