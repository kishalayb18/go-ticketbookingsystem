package main

import (
	"fmt"
	"strings"
)

const (
	conferenceName         = "Coldplay Coference"
	conferenceTotalTickets = 50
)

var (
	remainingTickets = 50
	bookingList      []string
)

func main() {

	var (
		userFirstName string
		userLastName  string
		userMail      string
		userTickets   int
		bookRow       int
	)

	greetMsg(remainingTickets)

	// ticket booking logic
	for {

		userFirstName, userLastName, userMail, userTickets, bookRow = getBookingDetails()
		isValidName, isValidMail, isValidTicketNumber, isValidRow := validateBooking(userFirstName, userLastName, userMail, userTickets, bookRow)

		if !isValidName || !isValidMail {
			fmt.Println("Please Provide Valid Name, Mail")
			continue
		} else if !isValidRow {
			fmt.Println("Not a valid choice for row")
			continue
		} else if !isValidTicketNumber {
			// ticket number validation logic
			tryValue := 1
			for {
				fmt.Printf("The number of tickets you could choose is 1 to %v\n", remainingTickets)
				fmt.Printf("Please enteer a valid number ")
				fmt.Scan(&userTickets)
				tryValue--
				if tryValue == 0 {
					fmt.Println("Not Valid Booking")
					break
				}
			}
		} else {
			bookingList, remainingTickets = ticktBooking(userFirstName, userLastName, userTickets, userMail, bookRow)
			greetMsg(remainingTickets)

		}
	}
}

func greetMsg(remainingTickets int) {
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Printf("Welcome to the %v Booking System\nWe have total %v tickets\n", conferenceName, conferenceTotalTickets)
	fmt.Printf("Remaining number of tickets %v \n", remainingTickets)
	fmt.Println("Get your tickets")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
}

// get details
func getBookingDetails() (string, string, string, int, int) {
	var (
		userFirstName string
		userLastName  string
		userMail      string
		userTickets   int
		bookRow       int
	)

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

	fmt.Println("Select the preferred row")
	fmt.Println("1. Front Row \t 2. Middle Row \t Back Row")
	fmt.Printf("Select the choice number ")
	fmt.Scan(&bookRow)
	fmt.Println("")

	return userFirstName, userLastName, userMail, userTickets, bookRow
}

// shortname method
func convertToShortName(fullName string) string {
	fullNmeArr := strings.Fields(fullName) // Split by space
	if len(fullNmeArr) < 2 {
		return fullNmeArr[0] // Return if only one name
	}
	initial := string(fullNmeArr[0][0]) // First letter of first name
	lastName := fullNmeArr[1]           // Last name
	return fmt.Sprintf("%s.%s", initial, lastName)
}

// validate datas logic
func validateBooking(userFirstName string, userLastName string, userMail string, userTickets int, bookRow int) (bool, bool, bool, bool) {
	isValidName := len(userFirstName) > 2 && len(userLastName) > 2
	isValidMail := strings.Contains(userMail, "@")
	isValidTicketNumber := userTickets <= remainingTickets || userTickets > 0
	isValidRow := bookRow >= 1 || bookRow <= 3

	return isValidName, isValidMail, isValidTicketNumber, isValidRow
}

// ticket booking
func ticktBooking(userFirstName string, userLastName string, userTickets int, userMail string, bookRow int) ([]string, int) {
	totalTicketBill, selectedRow := calculateBill(bookRow, userTickets)
	fmt.Printf("Booking Successful %v %v has booked %v %vtickets\n", userFirstName, userLastName, userTickets, selectedRow)
	fmt.Println("")

	fmt.Printf("Tickets have been sent to %v", userMail)
	fmt.Println("")
	fmt.Printf("Total bill %v", totalTicketBill)
	fmt.Println("")

	// update the booking list
	fullName := strings.ToUpper(userFirstName) + " " + strings.ToUpper(userLastName)
	bookingList = append(bookingList, convertToShortName(fullName))
	fmt.Println("")
	fmt.Printf("BookingList %v \n", bookingList)

	// remaining tickets
	remainingTickets = remainingTickets - userTickets

	return bookingList, remainingTickets

}

func calculateBill(bookRow int, userTickets int) (int, string) {
	var totalBill int
	var rowName string
	switch bookRow {
	case 1:
		totalBill = 5000 * userTickets
		rowName = "Front Row"
	case 2:
		totalBill = 3500 * userTickets
		rowName = "Middle Row"
	case 3:
		totalBill = 1500 * userTickets
		rowName = "back Row"
	default:
		fmt.Println("not a valid choice for row")
		totalBill = 0
		rowName = ""
	}
	return totalBill, rowName
}
