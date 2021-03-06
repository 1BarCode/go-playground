package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	// sb := spanishBot{}

	printGreeting(eb)
	// printGreeting(sb)
}


// interface is used to make printGreeting reusable for different argument type
func printGreeting(eb englishBot) {
	fmt.Println((eb.getGreeting()))
}

// func printGreeting(sb spanishBot) {
// 	fmt.Println((sb.getGreeting()))
// }

func (englishBot) getGreeting() string {
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}