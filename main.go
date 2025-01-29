package main

import (
	"fmt"
)

func main() {
	const (
		conferenceName         = "Coldplay Coference"
		conferenceTotalTickets = 50
		firstRowTicketPrice    = 5000
		midRowTicketPrice      = 3000
		lastRowTicketPrice     = 5000
	)

	var (
		remainingTickets = 50
		userFirstName    string
		userLastName     string
		userMail         string
		userTickets      int
		bookingList      []string

		arr_p_bl = 0
	)

	fmt.Printf("Welcome to the %v Booking System\nWe have total %v tickets\n", conferenceName, conferenceTotalTickets)
	fmt.Printf("Remaining number of tickets %v \n", remainingTickets)
	fmt.Println("Get your tickets")

	// userName = "Pat"

	fmt.Printf("Enter first name ")
	fmt.Scan(&userFirstName)
	fmt.Println("")

	fmt.Printf("Enter last name ")
	fmt.Scan(&userLastName)
	fmt.Println("")

	fmt.Printf("Enter mail ")
	fmt.Scan(&userMail)
	fmt.Println("")

	fmt.Printf("Enter number of ticket ")
	fmt.Scan(&userTickets)
	fmt.Println("")

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Booking Successful %v %v has booked %v tickets\n", userFirstName, userLastName, userTickets)
	fmt.Printf("Tickets have been sent to %v", userMail)

	bookingList = append(bookingList, userFirstName+" "+userLastName)
	arr_p_bl++
	fmt.Println(arr_p_bl)
	fmt.Printf("List %v \n", bookingList)

	fmt.Printf("Number of Remaining Tickets %v \n", remainingTickets)

}
