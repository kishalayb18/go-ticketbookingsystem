package process

import (
	"fmt"

	"github.com/kishalayb18/go-ticketbookingsystem/vars"
)

func PostBookingDetails(bookingCode string) {
	if vars.TicketCodesList[bookingCode] { // this is mapped with bool
		fullName, mail, selectedRow, numberofTickets, totalBill := GetBookingDetails(bookingCode)
		fmt.Printf("Name: %v \n", fullName)
		fmt.Printf("Mail: %v \n", mail)
		fmt.Printf("Selected row: %v\n", selectedRow)
		fmt.Printf("Total tickets: %v \n", numberofTickets)
		fmt.Printf("Total Bill: %v \n", totalBill)
		fmt.Println("")
	} else {
		fmt.Println("Booking Code is not correct")
	}
}
