package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// alex := person{"Alex", "Anderson"}
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)
	jim := person{firstName: "jim",
		lastName: "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 12345},
	}
	//&variable: give me the memory address of the value this variable is pointing at.
	//if jim is stored at RAM address 0001, jimPointer is directly pointing at the RAM address 0001
	//*pointer: Give me the value this memory address is pointing at
	//jimPointer is literally pointing to a memory address. it is NOT pointing to the value itself.
	jimPointer := &jim
	jimPointer.updateName("jimmy")
	jim.print()
}

/*
[0001][person{firstName:"Jim"}]
Turn address into value with *address
turn value into address with &value

When you want to modify a reference, pass in the memory address into the receiver. then the function will turn the address into a value.
A POINTER IS A POINTER TO A MEMORY ADDRESS.

It turns out that go will automatically convert a value to a pointer if the receiver is a pointer of a person
*/

//When you see a star next to a type, this is saying that it is
// a DESCRIPTION of a type. It means we're working with a pointer to a person.
// You shouldn't think of this as an operator. this is a type description like string, int, float64.
func (pointerToPerson *person) updateName(newFirstName string) {
	//star next to a value converts a pointer that points at a memory address to the actual struct value in the memory.
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
