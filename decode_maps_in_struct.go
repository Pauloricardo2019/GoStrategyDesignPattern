package main

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type Address struct {
	City  string
	State string
}

type Person struct {
	FirstName string
	LastName  string
	Address   Address
}

func main() {
	mapPerson := make(map[string]interface{})
	var person Person
	mapPerson["firstname"] = "Paulo"
	mapPerson["lastname"] = "Ricardo"
	mapAddress := make(map[string]interface{})
	mapAddress["city"] = "San Francisco"
	mapAddress["state"] = "California"
	mapPerson["address"] = mapAddress
	mapstructure.Decode(mapPerson, &person)

	test := person
	fmt.Println(test)

	fmt.Println(mapPerson)
}
