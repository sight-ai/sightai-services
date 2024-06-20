package comm_utils

import (
	"regexp"
)

func IsDateTimeValid(date string) bool {
	re := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
	return re.MatchString(date)
}

func IsStringIDValid(stringID string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9_.]*$")
	return re.MatchString(stringID)
}

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func IsEmailAddressValid(email string) bool {
	if len(email) < 3 && len(email) > 254 {
		return false
	}
	return emailRegex.MatchString(email)
}
