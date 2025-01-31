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
		bookRow          int
		totalBill        int
		bookingList      []string
	)

	greetMsg(conferenceName, conferenceTotalTickets, remainingTickets)

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

		fmt.Println("Select the preferred row")
		fmt.Println("1. Front Row \t 2. Middle Row \t Back Row")
		fmt.Printf("Select the choice number")
		fmt.Scan(&bookRow)
		fmt.Println("")

		// validation logic
		isValidName := len(userFirstName) > 2 && len(userLastName) > 2
		isValidMail := strings.Contains(userMail, "@")
		isValidTicketNumber := userTickets > remainingTickets || userTickets <= 0
		isValidRow := bookRow < 1 || bookRow > 3

		if !isValidName || !isValidMail {
			fmt.Println("Please Provide Valid Name, Mail")
			continue
		} else if !isValidRow {
			fmt.Println("Not a valid choice for row")
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
			switch bookRow {
			case 1:
				totalBill = 5000 * userTickets
			case 2:
				totalBill = 3500 * userTickets
			case 3:
				totalBill = 1500 * userTickets
			default:
				fmt.Println("not a valid choice for row")
			}
			fmt.Printf("Booking Successful %v %v has booked %v tickets\n", userFirstName, userLastName, userTickets)
			fmt.Println("")
			// remaining tickets
			remainingTickets = remainingTickets - userTickets

			fmt.Printf("Tickets have been sent to %v", userMail)
			fmt.Println("")

			// update the booking list
			fullName := strings.ToUpper(userFirstName) + " " + strings.ToUpper(userLastName)
			bookingList = append(bookingList, ConvertToShortName(fullName))
			fmt.Printf("BookingList %v \n", bookingList)
			fmt.Printf("Number of Remaining Tickets %v \n", remainingTickets)
			break

		}
	}
}

func greetMsg(conferenceName string, conferenceTotalTickets int, remainingTickets int) {
	fmt.Printf("Welcome to the %v Booking System\nWe have total %v tickets\n", conferenceName, conferenceTotalTickets)
	fmt.Printf("Remaining number of tickets %v \n", remainingTickets)
	fmt.Println("Get your tickets")
}

// shortname method
func ConvertToShortName(fullName string) string {
	fullNmeArr := strings.Fields(fullName) // Split by space
	if len(fullNmeArr) < 2 {
		return fullNmeArr[0] // Return if only one name
	}
	initial := string(fullNmeArr[0][0]) // First letter of first name
	lastName := fullNmeArr[1]           // Last name
	return fmt.Sprintf("%s.%s", initial, lastName)
}
