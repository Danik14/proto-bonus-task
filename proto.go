package main

import (
	"fmt"
	"test/gen"

	"github.com/golang/protobuf/ptypes/timestamp"
)

func main() {
	// Create a new PhoneNumber message.
	phone := &gen.PhoneNumber{
		Number: "123-456-7890",
		Type:   gen.PhoneType_MOBILE,
	}

	// Create a new Person message.
	person := &gen.Person{
		Name:        "John Doe",
		Id:          123,
		Email:       "johndoe@example.com",
		Phones:      []*gen.PhoneNumber{phone},
		LastUpdated: &timestamp.Timestamp{Seconds: 1619410882},
	}

	// Create a new AddressBook message.
	addressBook := &gen.AddressBook{
		People: []*gen.Person{person},
	}

	// Print the contents of the AddressBook message.
	fmt.Println(addressBook)
}
