package process

import (
	"fmt"
	"strings"

	"github.com/kishalayb18/go-ticketbookingsystem/helpers"
)

// ticket booking
func TicktBooking(userFirstName string, userLastName string, userTickets int, userMail string, bookRow int) ([]string, int) {
	totalTicketBill, selectedRow := calculateBill(bookRow, userTickets)
	fmt.Printf("Booking Successful %v %v has booked %v %vtickets\n", userFirstName, userLastName, userTickets, selectedRow)
	fmt.Println("")

	fmt.Printf("Tickets have been sent to %v", userMail)
	fmt.Println("")
	fmt.Printf("Total bill %v", totalTicketBill)
	fmt.Println("")

	// update the booking list
	fullName := strings.ToUpper(userFirstName) + " " + strings.ToUpper(userLastName)
	helpers.BookingList = append(helpers.BookingList, helpers.ConvertToShortName(fullName))
	fmt.Println("")
	fmt.Printf("BookingList %v \n", helpers.BookingList)

	// remaining tickets
	helpers.RemainingTickets = helpers.RemainingTickets - userTickets

	return helpers.BookingList, helpers.RemainingTickets

}
