package main

func main() {
	cards := newDeck()
	cards.print()
	//calling print function of custom type object
	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()

}
