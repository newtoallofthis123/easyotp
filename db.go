package easyotp

import (
	"database/sql"
	"fmt"
)

// Options: holds the options needed to work with EasyOtp
type Options struct {
	TableName string
	Pool      string
	OtpLen    int32
	hasInit   bool
}

// EasyOtp: The main struct that this package provides, contains all the context and connections
// necessary to carry out the otp functions
// New(conn *sql.DB) and NewWithOptions(conn *sql.DB, options Options) to initialized
type EasyOtp struct {
	conn    *sql.DB
	options Options
}

// New: Initializes the EasyOtp struct with the default options
func New(conn *sql.DB) EasyOtp {
	return EasyOtp{
		conn,
		DefaultOptions(),
	}
}

// NewWithOptions: Initializes the EasyOtp struct with the options passed in
func NewWithOptions(conn *sql.DB, options Options) EasyOtp {
	return EasyOtp{
		conn,
		options,
	}
}

// DefaultOptions: The Default options for the EasyOtp struct
func DefaultOptions() Options {
	return Options{
		"otp",
		NumericPool,
		8,
		false,
	}
}

// Init: Initializes a new Instance, with all the necessary changes to the database
// Warning: This makes permanent changes to the database
// This has to be used before using the opt functions
func (eo *EasyOtp) Init() error {
	otpQuery := fmt.Sprintf("CREATE TABLE if not exists %s(id VARCHAR(225), otp VARCHAR(225));", eo.options.TableName)
	_, err := eo.conn.Exec(otpQuery)
	eo.options.hasInit = true
	if err != nil {
		return err
	}
	return nil
}

func (eo *EasyOtp) addOtp(otp string) (string, error) {
	query := `
	INSERT INTO %s (id, otp)
	VALUES ($1, $2)
	`

	id := GenerateId()

	_, err := eo.conn.Exec(fmt.Sprintf(query, eo.options.TableName), id, otp)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (eo *EasyOtp) getOtp(id string) (string, error) {
	query := `
	SELECT * from %s WHERE id=$1
	`

	var otpId string
	var otp string
	row := eo.conn.QueryRow(fmt.Sprintf(query, eo.options.TableName), id)
	err := row.Scan(&otpId, &otp)
	if err != nil {
		return "", err
	}

	return otp, nil
}
