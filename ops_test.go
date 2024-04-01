package easyotp_test

import (
	"testing"

	"github.com/newtoallofthis123/easyotp"
)

func TestValidation(t *testing.T) {
	db := getDb()

	test := easyotp.New(db)
	err := test.Init()
	if err != nil {
		t.Error(err)
	}
	id, err := test.NewOtp("12345")
	if err != nil {
		t.Error(err)
	}

	correct, err := test.ValidateOtp(id, "12345")
	if err != nil {
		t.Error(err)
	}
	if !correct {
		t.Error("Validation fails")
	}
}
