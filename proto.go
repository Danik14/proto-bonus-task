package main

import (
	"fmt"
	"test/gen"

	"google.golang.org/protobuf/proto"
)

func main() {
	// Binary encoded phone number
	binData := []byte{10, 45, 10, 8, 74, 111, 104, 110, 32, 68, 111, 101, 16, 210, 9, 26, 16, 106, 100, 111, 101, 64, 101, 120, 97, 109, 112, 108, 101, 46, 99, 111, 109, 34, 12, 10, 8, 53, 53, 53, 45, 52, 51, 50, 49, 16, 3}

	// Unmarshal the binary data into a AddressBook message
	person := &gen.AddressBook{}
	err := proto.Unmarshal(binData, person)
	if err != nil {
		panic(err)
	}

	fmt.Println(person)
}
