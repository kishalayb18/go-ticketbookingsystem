package process

import (
	"fmt"

	"github.com/kishalayb18/go-ticketbookingsystem/helpers"
	"github.com/kishalayb18/go-ticketbookingsystem/vars"
)

func BookTicket() {
	var (
		userFirstName string
		userLastName  string
		userMail      string
		userTickets   int
		bookRow       int
	)

	// ticket booking logic
	for {

		userFirstName, userLastName, userMail, userTickets, bookRow = SetUserDetails()
		isValidName, isValidMail, isValidTicketNumber, isValidRow := helpers.ValidateBooking(userFirstName, userLastName, userMail, userTickets, bookRow)

		if !isValidName {
			fmt.Println("Please Provide Valid Name")
			continue
		} else if !isValidMail {
			fmt.Println("Please Provide Valid Mail")
			continue
		} else if !isValidRow {
			fmt.Println("Not a valid choice for row")
			continue
		} else if !isValidTicketNumber {
			// ticket number validation logic
			tryValue := 1
			for {
				fmt.Printf("The number of tickets you could choose is 1 to %v\n", vars.RemainingTickets)
				fmt.Printf("Please enteer a valid number ")
				fmt.Scan(&userTickets)
				tryValue--
				if tryValue == 0 {
					fmt.Println("Not Valid Booking")
					break
				}
			}
		} else {
			ValidTicktBooking(userFirstName, userLastName, userTickets, userMail, bookRow)
			break
		}
	}
}
