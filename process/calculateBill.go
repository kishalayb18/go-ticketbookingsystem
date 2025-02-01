package process

import "fmt"

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
