package process

import (
	"fmt"
	"strings"
)

// get details
func SetUserDetails() (string, string, string, int, int) {
	var (
		userFirstName string
		userLastName  string
		userMail      string
		userTickets   int
		bookRow       int
	)

	fmt.Printf("Enter first name ")
	fmt.Scan(&userFirstName)
	userFirstName = strings.ToUpper(userFirstName)
	fmt.Println("")

	fmt.Printf("Enter last name ")
	fmt.Scan(&userLastName)
	userLastName = strings.ToUpper(userLastName)
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
