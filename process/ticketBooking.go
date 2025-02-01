package process

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kishalayb18/go-ticketbookingsystem/helpers"
	"github.com/kishalayb18/go-ticketbookingsystem/vars"
)

// ticket booking
func TicktBooking(userFirstName string, userLastName string, userTickets int, userMail string, bookRow int) ([]string, int) {
	totalTicketBill, selectedRow := calculateBill(bookRow, userTickets)

	var userTicketDetails = make(map[string]string)
	userTicketDetails["firstName"] = userFirstName
	userTicketDetails["lastName"] = userLastName
	userTicketDetails["totalTickets"] = strconv.FormatInt(int64(userTickets), 10)
	userTicketDetails["mail"] = userMail
	userTicketDetails["selectedRow"] = selectedRow
	userTicketDetails["totalTickets"] = strconv.FormatInt(int64(totalTicketBill), 10)

	fmt.Printf("Booking Successful %v %v has booked %v %v tickets\n", strings.ToUpper(userFirstName), strings.ToUpper(userLastName), userTickets, selectedRow)
	fmt.Println("")

	fmt.Printf("Tickets have been sent to %v", userMail)
	fmt.Println("")
	fmt.Printf("Total bill %v", totalTicketBill)
	fmt.Println("")

	// update the booking list
	fullName := strings.ToUpper(userFirstName) + " " + strings.ToUpper(userLastName)
	vars.BookingList = append(vars.BookingList, helpers.ConvertToShortName(fullName))
	fmt.Println("")
	fmt.Printf("BookingList %v \n", vars.BookingList)

	// booking list details
	vars.BookingListDetails = append(vars.BookingListDetails, userTicketDetails)
	fmt.Println("Details of List ")
	fmt.Printf("%v", vars.BookingListDetails)

	// remaining tickets
	vars.RemainingTickets = vars.RemainingTickets - userTickets

	return vars.BookingList, vars.RemainingTickets

}
