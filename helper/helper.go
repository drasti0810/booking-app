package helper

func ValidateInput(userName string, userTickets uint, ticketCount uint) (bool, bool){
	isValidName := len(userName) >= 2
	isValidTicket := userTickets <= ticketCount && userTickets > 0
	return isValidName, isValidTicket
}