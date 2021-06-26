package routes

import "errors"

func validateNotEmptyString(s string) (bool, error) {
	if s == "" {
		return false, errors.New("Empty String")
	}
	return true, nil
}

func validateEmail(s string) (bool, error) {
	if s == "" {
		return false, errors.New("Empty E-mail")
	}
	return true, nil
}

func validatePassword(s string) (bool, error) {
	if s == "" {
		return false, errors.New("Empty Password")
	}
	return true, nil
}
