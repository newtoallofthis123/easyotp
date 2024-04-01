package easyotp

import (
	"math/rand"
)

// GenerateNumeric generates a random number of length len
func GenerateNumeric(len int) string {
	return Generate(len, NumericPool)
}

// GenerateAlphaNumeric: Generates an otp from the AlphaNumeric Pool
func GenerateAlphaNumeric(len int) string {
	return Generate(len, AlphaNumericPool)
}

// Generate: Generates a random string with a given length out of a given pool
func Generate(length int, pool string) string {
	var otp string

	for i := 0; i < length; i++ {
		otp += string(pool[rand.Intn(len(pool))])
	}

	return otp
}

// AlphaNumericPool: Pool that has A-Z, a-z and 0-9
const AlphaNumericPool = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// NumericPool: Pool with 0-9
const NumericPool = "0123456789"

// Generated a random ID for the otp
func GenerateId() string {
	return GenerateAlphaNumeric(32)
}
