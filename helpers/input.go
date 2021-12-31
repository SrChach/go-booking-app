package helpers

import "fmt"

func GetInputString(inputMessage string) string {
	var inputString string
	fmt.Println(inputMessage)
	fmt.Scan(&inputString)

	return inputString
}

func GetClientName() string {
	return GetInputString("Please enter your name")
}

func GetClientLastName() string {
	return GetInputString("Please enter your last name")
}

func GetNumberOfTickets() uint {
	var ticketsNumber uint

	fmt.Println("Please enter the number of tickets you want")
	fmt.Scan(&ticketsNumber)

	return ticketsNumber
}

/** Multiple Returns are allowed in Go! */
func GetUserBookingData() (string, string, uint) {
	return GetClientName(), GetClientLastName(), GetNumberOfTickets()
}
