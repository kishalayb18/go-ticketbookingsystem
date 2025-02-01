package helpers

import (
	"fmt"

	"github.com/kishalayb18/go-ticketbookingsystem/vars"
)

func GreetMsg(RemainingTickets int) {
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Printf("Welcome to the %v Booking System\nWe have total %v tickets\n", vars.ConferenceName, vars.ConferenceTotalTickets)
	fmt.Printf("Remaining number of tickets %v \n", RemainingTickets)
	fmt.Println("Get your tickets")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
}
