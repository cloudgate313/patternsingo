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

//  *******  contactInfo == contactInfo: contactInfo *******
// type person struct {
// 	firstname string
// 	lastName  string
// 	contactInfo
// }

// type person struct {
// 	firstname string
// 	lastName  string
// 	contactInfo: contactInfo
// }

func main() {
	var anthony person

	// dot notation
	anthony.firstname = "Anthony"
	anthony.lastName = "Evans"
	// fmt.Println(anthony)

	// object* notation
	evan := person{
		firstname: "Evan",
		lastName:  "Armbiser",
		contact: contactInfo{
			email:   "evananthony13@gmail.com",
			zipcode: 10010,
		},
	}
	// fmt.Printf("%+v", evan)

	evan.updateFirstName("Anthony")
	evan.print()
}

func (p *person) updateFirstName(updatedName string) {
	p.firstname = updatedName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
