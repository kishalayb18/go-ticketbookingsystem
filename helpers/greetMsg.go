package helpers

import "fmt"

func GreetMsg(RemainingTickets int) {
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Printf("Welcome to the %v Booking System\nWe have total %v tickets\n", ConferenceName, ConferenceTotalTickets)
	fmt.Printf("Remaining number of tickets %v \n", RemainingTickets)
	fmt.Println("Get your tickets")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
}
