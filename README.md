# EasyOTP

The OTP system for the quick engineer

OTP's are one of the most easiest, yet easy to mess up part of web development.
I have seen many engineers using even the local storage for otp verification.
All of this leads to security issues and more.

EasyOTP is a very easy way to implement an OTP system in Go
Here how easy it is:

```go
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/newtoallofthis123/easyotp"
)

func main() {
	db, _ := sql.Open("sqlite3", "test.db")

	otp := easyotp.New(db)
	otp.Init()

	 id, _ := otp.NewOtp(easyotp.GenerateNumeric(6))
	id, _ := otp.NewOtp("12345") // We are using 12345 for the sake of convience

	correct, _ := otp.ValidateOtp(id, "12345")
	if correct {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
```

It is that easy!

But, here is a small doc of all the stuff you may need, so consider this the mini doc :)

## Mini Doc

### General Functions

- `New(conn *sql.DB) EasyOtp`: Returns an instance to the EasyOtp struct with all the default options 

- `NewWithOptions(conn *sql.DB, options Options) EasyOtp`: Returns an instance to the EasyOtp struct with the passed in options struct

- `Options Struct`: 
```go
type Options struct {
	TableName string
	Pool      string
	OtpLen    int32
    hasInit   bool  This cannot be changed or accessed
}
```

The default options are as follows:

```go
Options{
	"otp",
	NumericPool,
	8,
    false,
}
```

Hence, the default table name would be `otp`, the otp would be generated from the NumericPool of data and the length
would be 8.
You can choose from quite some pools to generate the otp from, but the pool can be pretty much any string.

- `const AlphanumericPool`: Has A-Z and 0-9
- `const NumericPool`: Has 0-9

- `GenerateNumeric(len int) string`: Generates a otp with the given length using the numeric pool
- `GenerateAlphaNumeric(len int) string`: Generates a otp with the given length using the AlphaNumeric pool
- `Generate(length int, pool string) stirng`: Generates a otp / random string of a given length using the given pool

### Instance Functions

- `Init() error`: Initializes the database connection to the db conn provided and writes the table to the db. This makes permanent changes to the database.
Any and all instances functions can only be performed after this initialization

- `NewOtp(otp string) (string, error)`: Creates a new Otp entry in the database and returns the id.
This id can be passed into the context of any web server or application that needs to authenticate using otp
Otp has to be passed in manually, or using the many generators provided 
also, this function can only be used with an instance

`ValidateOtp(id string, otp string) (bool, error)`: Validates the provided otp and the id and returns if it is correct or not Uses a simple select query on the database.
Needs the context of the instance as well.
This can be used by storing the id in the localstorage or the cookies of the browser,
then getting it in the auth request, then validating the otp like this

## License

The above project is licensed under the MIT License, the details of which can be found in the [LICENSE](LICENSE) file
