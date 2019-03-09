package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode int
}

type person struct {
	firstname string
	lastName  string
	contact   contactInfo
}

func main() {
	evan := person{
		firstname: "Evan",
		lastName:  "Armbiser",
		contact: contactInfo{
			email:   "evananthony13@gmail.com",
			zipcode: 10010,
		},
	}

	pointerToEvan := &evan

	evan.print()
	pointerToEvan.updateFirstName("Anthony")
	pointerToEvan.print()
}

func (pointerToPerson *person) updateFirstName(updatedName string) {
	(*pointerToPerson).firstname = updatedName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
