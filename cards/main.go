package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
	//calling print function of custom type object
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()
	//Calling toString function on deck of cards
	// cards1 := newDeck()
	// cards1.saveToFile("my_saved_file")
	// cards2 := newDeckFromFile("my_saved_file")
	// cards2.print()
}
