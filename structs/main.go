package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName   string
	lastName    string
	contactInfo contactInfo
}

func main() {
	// // alex :=  person{"Alex", "Anderson"}
	// // alex := person{firstName: "Alex", lastName: "Anderson"}
	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderon"
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 9400,
		},
	}

	jim.updateName("jimmy")
	jim.print()
	// fmt.Printf("%+v", jim)
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
