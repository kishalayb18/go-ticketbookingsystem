package process

import (
	"fmt"
	"sync"
	"time"
)

func sendTicket(ticket string, totalTicketBill int, bookingCode string, wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(15 * time.Second)
	fmt.Println(ticket)
	fmt.Printf("Total bill: %v \n", totalTicketBill)
	fmt.Printf("Booking code: %v \n", bookingCode)
	fmt.Println("")
}
