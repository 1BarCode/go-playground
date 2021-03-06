package main

import "fmt"

// type contactInfo struct {
// 	email string
// 	zipCode int
// }

// type person struct {
// 	firstName string
// 	lastName  string
// 	contactInfo // or contact contactInfo
// }

// func main() {
// 	jim := person{
// 		firstName: "Jim",
// 		lastName: "Party",
// 		contactInfo: contactInfo{
// 			email: "jim@gmail.com",
// 			zipCode: 94000,
// 		},
// 	}

// 	jim.updateName("jimmy")
// 	jim.print()
// }

// func (pointerToPerson *person) updateName(newFirstName string) {
// 	(*pointerToPerson).firstName = newFirstName
// }

// func (p person) print() {
// 	fmt.Printf("%+v", p)
// }

// ============ test

func main() {
	name := "bill"

	namePointer := &name

	fmt.Println(&name)
	fmt.Println(namePointer)
	printPointer(namePointer)
}

func printPointer(namePointer *string) {
	fmt.Println(&namePointer)
}
