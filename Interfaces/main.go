package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}

type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb) //Gives error if we try to use function overloading in Go

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sb spanishBot) {
// 	//Same logic, just the different type in arguments, Go does not support function overloading
// 	//Hence Interfaces becomes handy in Go
// 	fmt.Println(sb.getGreeting())
// }

func (eb englishBot) getGreeting() string {
	//Very custom logic for generating english greeting
	return "Hello world!"
}

func (sb spanishBot) getGreeting() string {
	//Very custom logic for generating spanish greeting
	return "Hola!"
}
