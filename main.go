package main

import (
	"fmt"
	"sync"

	"github.com/kishalayb18/go-ticketbookingsystem/helpers"
	"github.com/kishalayb18/go-ticketbookingsystem/vars"
)

var wg = sync.WaitGroup{}

func main() {
	var (
		choice int
	)

	defer wg.Wait()

	for {
		if vars.RemainingTickets <= 0 {
			fmt.Println("")
			fmt.Println("---------------------------------------------")
			fmt.Println("House Full")
			fmt.Println("New booking has stopped")
			fmt.Println("---------------------------------------------")
			fmt.Println("Concert Menu")
			fmt.Println("---------------------------------------------")
			fmt.Println("Choose From The Menu")
			fmt.Println("1. View Your Ticket Details 2. View Booking List")
			fmt.Println("")
			fmt.Printf("Select the choice ")
			fmt.Scan(&choice)
			fmt.Println("---------------------------------------------")
			fmt.Println("")
			choice = choice + 1
			menuChoice(choice)
		} else {
			fmt.Println("")
			fmt.Println("---------------------------------------------")
			helpers.GreetMsg(vars.RemainingTickets)
			fmt.Println("Concert Menu")
			fmt.Println("---------------------------------------------")
			fmt.Println("Choose From The Menu")
			fmt.Println("1. Book Ticket \t 2. View Your Ticket Details 3. View Booking List")
			fmt.Println("")
			fmt.Printf("Select the choice ")
			fmt.Scan(&choice)
			fmt.Println("---------------------------------------------")
			fmt.Println("")
			menuChoice(choice)
		}
	}
}
