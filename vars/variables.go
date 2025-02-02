package vars

var (
	RemainingTickets = 50

	// lists
	BookingList []string
	// BookingListDetails []BookedTicketData

	// maps
	TicketCodesList   = make(map[string]bool)
	BookingDetailsMap = make(map[string]BookedTicketData)
)
