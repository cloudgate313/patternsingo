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

	evan.updateFirstName("Anthony")
	evan.print()
}

func (p *person) updateFirstName(updatedName string) {
	(*p).firstname = updatedName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
