package helpers

import (
	"fmt"
	"strings"
)

// shortname method
func ConvertToShortName(fullName string) string {
	fullNmeArr := strings.Fields(fullName) // Split by space
	if len(fullNmeArr) < 2 {
		return fullNmeArr[0] // Return if only one name
	}
	initial := string(fullNmeArr[0][0]) // First letter of first name
	lastName := fullNmeArr[1]           // Last name
	return fmt.Sprintf("%s.%s", initial, lastName)
}
