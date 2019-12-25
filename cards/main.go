package main

import "fmt"

func newCard() string {
	return "Five of Spades"
}

func main() {
	cards := []string{"Ace of heart", newCard()}
	// append function does not modify existing slice,
	//instead it returns a new slice and assigns back to cards slice
	cards = append(cards, "Joker!")
	//Let's iterate over slice
	for i, card := range cards {
		fmt.Println(i, card)
	}
}
