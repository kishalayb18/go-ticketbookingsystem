package process

import (
	"fmt"

	"github.com/kishalayb18/go-ticketbookingsystem/helpers"
	"github.com/kishalayb18/go-ticketbookingsystem/tickets"
	"github.com/kishalayb18/go-ticketbookingsystem/vars"
)

func ValidTicktBooking(userFirstName string, userLastName string, userTickets int, userMail string, bookRow int) {

	bookingCode := tickets.GenerateTicketCode()
	totalTicketBill, selectedRow := calculateBill(userTickets, bookRow)

	var bookedTicketData = vars.BookedTicketData{
		FirstName:   userFirstName,
		LastName:    userLastName,
		Mail:        userMail,
		Tickets:     userTickets,
		SelectedRow: selectedRow,
		BookingCode: bookingCode,
		Bill:        totalTicketBill,
	}
	// booking list details
	vars.BookingDetailsMap[bookingCode] = bookedTicketData

	// update the booking name list
	fullName := userFirstName + " " + userLastName
	vars.BookingList = append(vars.BookingList, helpers.ConvertToShortName(fullName))

	// message
	fmt.Printf("Booking Successful %v %v has booked %v %v tickets\n", userFirstName, userLastName, userTickets, selectedRow)
	fmt.Printf("Tickets have been sent to %v \n", userMail)
	fmt.Printf("Total bill %v \n", totalTicketBill)
	fmt.Printf("Booking code %v \n", bookingCode)
	fmt.Println("")

	// fmt.Println("list")
	// fmt.Printf("%v", vars.TicketCodesList)

	// remaining tickets
	vars.RemainingTickets = vars.RemainingTickets - userTickets

	// return vars.BookingList, vars.RemainingTickets

}
