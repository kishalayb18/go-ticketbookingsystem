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
	)

	fmt.Printf("Welcome to the %v Booking System\nWe have total %v tickets\n", conferenceName, conferenceTotalTickets)
	fmt.Printf("Remaining number of tickets %v \n", remainingTickets)
	fmt.Println("Get your tickets")

	// ticket booking logic
	for {

		// user details fnmae
		fmt.Printf("Enter first name ")
		fmt.Scan(&userFirstName)
		fmt.Println("")

		// user details lname
		fmt.Printf("Enter last name ")
		fmt.Scan(&userLastName)
		fmt.Println("")

		// user details mail
		fmt.Printf("Enter mail ")
		fmt.Scan(&userMail)
		fmt.Println("")

		// user details tickets number
		fmt.Printf("Enter number of ticket ")
		fmt.Scan(&userTickets)

		//name validation logic
		isValidName := len(userFirstName) > 2 && len(userLastName) > 2
		isValidMail := strings.Contains(userMail, "@")
		isValidTicketNumber := userTickets > remainingTickets || userTickets <= 0

		if !isValidName || !isValidMail {
			fmt.Println("Please Provide Valid Name, Mail")
			continue
		} else if !isValidTicketNumber {
			// ticket number validation logic
			tryValue := 2
			for {
				fmt.Printf("The numver of tickets you could choose is 1 to %v\n", remainingTickets)
				// fmt.Printf("Number of Remaining Tickets %v \n", remainingTickets)
				fmt.Println("")
				fmt.Printf("Please enteer a valid number ")
				fmt.Scan(&userTickets)
				tryValue--
				if tryValue == 0 {
					fmt.Println("Not Valid Booking")
					break
				}
			}
		} else {
			// valid booking
			fmt.Println("")
			fmt.Printf("Booking Successful %v %v has booked %v tickets\n", userFirstName, userLastName, userTickets)
			fmt.Println("")
			// remaining tickets
			remainingTickets = remainingTickets - userTickets

			fmt.Printf("Tickets have been sent to %v", userMail)
			fmt.Println("")

			// update the booking list
			fullName := userFirstName + " " + userLastName
			shortName := ConvertToShortName(fullName)

			bookingList = append(bookingList, shortName)
			fmt.Printf("BookingList %v \n", bookingList)
			fmt.Printf("Number of Remaining Tickets %v \n", remainingTickets)
			break

		}

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
