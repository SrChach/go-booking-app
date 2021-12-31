package helpers

import (
	"fmt"
	"strconv"
)

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

/** Using a Map to store data */
func CreateBookingMap() map[string]string {
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
