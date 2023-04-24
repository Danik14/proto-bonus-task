package main

import (
	"fmt"
	"test/gen"

	"github.com/golang/protobuf/ptypes/timestamp"
)

func main() {
	phone := &gen.PhoneNumber{
		Number: "123-456-7890",
		Type:   gen.PhoneType_MOBILE,
	}

	person := &gen.Person{
		Name:        "Ryan Gosling",
		Id:          1,
		Email:       "ryan@gosling.com",
		Phones:      []*gen.PhoneNumber{phone},
		LastUpdated: &timestamp.Timestamp{Seconds: 12345},
	}

	addressBook := &gen.AddressBook{
		People: []*gen.Person{person},
	}

	fmt.Println(addressBook)
}
