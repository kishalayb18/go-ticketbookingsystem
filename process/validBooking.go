package process

import (
	"fmt"
	"sync"

	"github.com/kishalayb18/go-ticketbookingsystem/helpers"
	"github.com/kishalayb18/go-ticketbookingsystem/tickets"
	"github.com/kishalayb18/go-ticketbookingsystem/vars"
)

func ValidTicktBooking(userFirstName string, userLastName string, userTickets int, userMail string, bookRow int, wg *sync.WaitGroup) {

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

	//
	fmt.Println("Booking Successful ")
	var ticket = fmt.Sprintf("Confirm Ticket %v %v has booked %v %v tickets\n", userFirstName, userLastName, userTickets, selectedRow)
	fmt.Printf("Tickets have been sent to %v \n", userMail)

	//multi threading, wait group
	wg.Add(1)
	go sendTicket(ticket, totalTicketBill, bookingCode, wg)

	// fmt.Printf("%v", vars.TicketCodesList)
	// remaining tickets
	vars.RemainingTickets = vars.RemainingTickets - userTickets

}
