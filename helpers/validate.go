package helpers

import (
	"strings"

	"github.com/kishalayb18/go-ticketbookingsystem/vars"
)

// validate datas logic
func ValidateBooking(userFirstName string, userLastName string, userMail string, userTickets int, bookRow int) (bool, bool, bool, bool) {
	isValidName := len(userFirstName) > 2 && len(userLastName) > 2
	isValidMail := strings.Contains(userMail, "@")
	isValidTicketNumber := userTickets <= vars.RemainingTickets || userTickets > 0
	isValidRow := bookRow >= 1 || bookRow <= 3

	return isValidName, isValidMail, isValidTicketNumber, isValidRow

}
