package main

import (
	"fmt"

	"github.com/kishalayb18/go-ticketbookingsystem/process"
	"github.com/kishalayb18/go-ticketbookingsystem/vars"
)

func menuChoice(choice int) {
	switch choice {
	case 1:
		process.BookTicket(&wg)
	case 2:
		var bookingCode string
		fmt.Printf("Enter Your Booking Code ")
		fmt.Scan(&bookingCode)
		process.PostBookingDetails(bookingCode)
	case 3:
		fmt.Println("Booking List")
		for r, shortName := range vars.BookingList {
			fmt.Printf("%v. %v", r+1, shortName)
			fmt.Println("")
		}
	default:
		fmt.Println("wrong choice")
	}
}
