package booking

import (
	"fmt"
	"strconv"
)

func getInputString(inputMessage string) string {
	var inputString string
	fmt.Println(inputMessage)
	fmt.Scan(&inputString)

	return inputString
}

func getClientName() string {
	return getInputString("Please enter your name")
}

func getClientLastName() string {
	return getInputString("Please enter your last name")
}

func getNumberOfTickets() uint {
	var ticketsNumber uint

	fmt.Println("Please enter the number of tickets you want")
	fmt.Scan(&ticketsNumber)

	return ticketsNumber
}

/** Multiple Returns are allowed in Go! */
func GetUserBookingData() (string, string, uint) {
	return getClientName(), getClientLastName(), getNumberOfTickets()
}

/** using and returning a structure */
func GetBookingStruct() BookingData {
	clientName, clientLastName, numberOfTickets := GetUserBookingData()

	var bookingData = BookingData{
		Name:     clientName,
		LastName: clientLastName,
		Quantity: numberOfTickets,
	}

	return bookingData
}

/** Example of using a Map to store data */
func createBookingMap() map[string]string {
	clientName, clientLastName, numberOfTickets := GetUserBookingData()
	/** Making a Map
	make(map[string]string)
	*/
	stringNumberOfTickets := strconv.FormatUint(uint64(numberOfTickets), 10)
	userBookingData := map[string]string{
		"name":     clientName,
		"lastName": clientLastName,
		"quantity": stringNumberOfTickets,
	}
	return userBookingData
}
