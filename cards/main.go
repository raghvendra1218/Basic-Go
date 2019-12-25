package main

func newCard() string {
	return "Five of Spades"
}

func main() {
	cards := deck{"Ace of heart", newCard()}
	// append function does not modify existing slice,
	//instead it returns a new slice and assigns back to cards slice
	cards = append(cards, "Joker!")
	//calling print function of custom type object
	deck.print(cards)

}
