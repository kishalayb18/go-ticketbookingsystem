package process

import "github.com/kishalayb18/go-ticketbookingsystem/vars"

func GetBookingDetails(bookingCode string) (string, string, string, int, int) {
	firstname := vars.BookingDetailsMap[bookingCode].FirstName
	lastName := vars.BookingDetailsMap[bookingCode].LastName
	fullName := firstname + " " + lastName
	mail := vars.BookingDetailsMap[bookingCode].Mail
	selectedRow := vars.BookingDetailsMap[bookingCode].SelectedRow
	numberofTickets := vars.BookingDetailsMap[bookingCode].Tickets
	totalBill := vars.BookingDetailsMap[bookingCode].Tickets

	return fullName, mail, selectedRow, numberofTickets, totalBill
}
