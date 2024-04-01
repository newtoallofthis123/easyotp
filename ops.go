package easyotp

import "errors"

// NewOtp: Creates a new Otp entry in the database and returns the id.
// This id can be passed into the context of any web server or application that needs to
// authenticate using otp
// Otp has to be passed in manually, or using the many generators provided
// also, this function can only be used with an instance
func (eo *EasyOtp) NewOtp(otp string) (string, error) {
	if !eo.options.hasInit {
		return "", errors.New("Instance has not been initialized. Call the Init() function to initialize")
	}
	id, err := eo.addOtp(otp)
	if err != nil {
		return "", err
	}

	return id, nil
}

// ValidateOtp: Validates the provided otp and the id and returns if it is correct or not
// Uses a simple select query on the database.
// Needs the context of the instance as well.
// This can be used by storing the id in the localstorage or the cookies of the browser,
// then getting it in the auth request, then validating the otp like this
func (eo *EasyOtp) ValidateOtp(id string, otp string) (bool, error) {
	if !eo.options.hasInit {
		return false, errors.New("Instance has not been initialized. Call the Init() function to initialize")
	}
	actualOtp, err := eo.getOtp(id)
	if err != nil {
		return false, err
	}

	if otp == actualOtp {
		return true, nil
	} else {
		return false, nil
	}
}
