package main

import (
	"fmt"
	"strings"
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

		// arr_p_bl = 0
	)

	fmt.Printf("Welcome to the %v Booking System\nWe have total %v tickets\n", conferenceName, conferenceTotalTickets)
	fmt.Printf("Remaining number of tickets %v \n", remainingTickets)
	fmt.Println("Get your tickets")

	// userName = "Pat"

	// ticket booking logic
	for {

		// user details
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

		// remaining tickets
		remainingTickets = remainingTickets - userTickets

		fmt.Printf("Booking Successful %v %v has booked %v tickets\n", userFirstName, userLastName, userTickets)
		fmt.Printf("Tickets have been sent to %v", userMail)

		fullName := userFirstName + " " + userLastName

		shortName := ConvertToShortName(fullName)

		// update the booking list

		bookingList = append(bookingList, shortName)

		fmt.Printf("BookingList %v \n", bookingList)

		fmt.Printf("Number of Remaining Tickets %v \n", remainingTickets)
	}

}

// shortname method
func ConvertToShortName(fullName string) string {
	fullNmeArr := strings.Fields(fullName) // Split by space
	if len(fullNmeArr) < 2 {
		return strings.ToUpper(fullNmeArr[0]) // Return if only one name
	}
	initial := strings.ToUpper(string(fullNmeArr[0][0])) // First letter of first name
	lastName := strings.ToUpper(fullNmeArr[1])           // Last name
	return fmt.Sprintf("%s.%s", initial, lastName)
}
