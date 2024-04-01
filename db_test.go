package easyotp_test

import (
	"database/sql"
	"os"
	"reflect"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/newtoallofthis123/easyotp"
)

func getDb() *sql.DB {
	os.Remove("test.db")

	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

func TestNew(t *testing.T) {
	db := getDb()

	test := easyotp.New(db)
	if reflect.TypeOf(test) != reflect.TypeOf(easyotp.EasyOtp{}) {
		t.Errorf("New failed!")
	}

	defer db.Close()
}

func TestNewWithOptions(t *testing.T) {
	db := getDb()

	test := easyotp.NewWithOptions(db, easyotp.DefaultOptions())
	if reflect.TypeOf(test) != reflect.TypeOf(easyotp.EasyOtp{}) {
		t.Errorf("NewWithOptions failed!")
	}

	defer db.Close()
}

func TestInit(t *testing.T) {
	db := getDb()

	test := easyotp.NewWithOptions(db, easyotp.DefaultOptions())
	err := test.Init()
	if err != nil {
		t.Error(err)
	}

	defer db.Close()
}
