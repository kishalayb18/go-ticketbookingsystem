package process

import "fmt"

func calculateBill(userTickets int, bookRow int) (int, string) {
	var totalBill int
	var selectedRow string
	switch bookRow {
	case 1:
		totalBill = 5000 * userTickets
		selectedRow = "Front Row"
	case 2:
		totalBill = 3500 * userTickets
		selectedRow = "Middle Row"
	case 3:
		totalBill = 1500 * userTickets
		selectedRow = "back Row"
	default:
		fmt.Println("not a valid choice for row")
		totalBill = 0
		selectedRow = ""
	}
	return totalBill, selectedRow
}
