package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// Option -1: Declaring and assigning ,order in which they are declared they are assigned
	// Raghav := person{"Raghav", "Dixit"}
	// Option -2: Declaring and assigning, explicitly assigning the variables
	// Raghav := person{firstName: "Raghav", lastName: "Dixit"}
	//Option -3: Declare a person variable first and then assign the values
	// var Raghav person
	// fmt.Println(Raghav)
	// Raghav.firstName = "Raghav"
	// Raghav.lastName = "Dixit"
	//Prints out the struct variables along with the values
	// fmt.Printf("%+v", Raghav)
	//Embed one struct into another
	jim := person{
		firstName: "Jim",
		lastName:  "Anderson",
		contact: contactInfo{
			email:   "jim.aderson@gmail.com",
			zipCode: 95112,
		},
	}

	//Remember : Golang is pass by value language, that means whenever we pass value in arguments as
	//a receiver or as a function it will always create a copy of that value and function or receiver
	//will always work on the copy of that instance, Hence pointers are handy whenever we have to work
	//on pass by reference types or in other words work on the actual instance of object/struct/variable.

	jim.print()
	pointerToJim := &jim
	pointerToJim.updateName("Johnson")
	//Go allows receiver function to operate with person type and pointer to person as well
	//Hence we don't have to get the reference everytime for the person pointer and then call updateName
	//Above two lines can be replaced with one line below
	jim.updateName("Johnson1")
	jim.print()
}

//Create Receiver function for structs types
func (p person) print() {
	fmt.Printf("%+v", p)
}

//Update the firstName of the person
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
