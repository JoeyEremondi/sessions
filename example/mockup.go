package main

import "github.com/JoeyEremondi/GoSesh/mockup"

// ASendToBThenBToC : A sends an int message to B, B sends the message to C
func ASendToBThenBToC() {

	channelAToB := mockup.Channel{
		Name:        "fromAtoB",
		Source:      "A",
		Destination: "B"}

	message := mockup.MessageType{Type: "int"}
	sendAToB := mockup.Send(channelAToB, message)

	channelBToC := mockup.Channel{
		Name:        "fromBtoC",
		Source:      "B",
		Destination: "C"}

	sendBToC := mockup.Send(channelBToC, message)

	mockup.CreateStubProgram("ABCExample",
		sendAToB,
		sendBToC,
	)
}

func main() {
	ASendToBThenBToC()
}
