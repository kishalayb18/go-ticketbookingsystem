package process

import (
	"fmt"
	"time"
)

func sendTicket(ticket string, totalTicketBill int, bookingCode string) {
	time.Sleep(15 * time.Second)
	fmt.Println(ticket)
	fmt.Printf("Total bill: %v \n", totalTicketBill)
	fmt.Printf("Booking code: %v \n", bookingCode)
	fmt.Println("")
}
