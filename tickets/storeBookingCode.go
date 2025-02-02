package tickets

import (
	"github.com/kishalayb18/go-ticketbookingsystem/vars"
)

func storeCode(code string) bool {

	if vars.TicketCodesList[code] { // if this is true that means the code is present
		newCode := GenerateTicketCode() // need to generate another code
		storeCode(newCode)              // recall the function to validate the same
		return false
	} else {
		vars.TicketCodesList[code] = true
		return true
	}
}
