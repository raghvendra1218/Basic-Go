package main

import "fmt"

func newCard() string {
	return "Five of Spades"
}

func main() {
	card := newCard()
	fmt.Println(card)
}
