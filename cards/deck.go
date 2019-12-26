package main

import "fmt"

//Create a new type of deck,
//which is a slice of strings
type deck []string

//Create a new Deck, we do not want a receiver function, as we need new deck, we are not working
//on an already created deck
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "King", "Queen", "Jack"}
	//loop through suit and values to create a deck of card
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			newCard := value + " of " + suit
			cards = append(cards, newCard)
		}
	}
	return cards
}

//Deal function to return decks of hand and remaining cards
//This is also an example of multiple return values from a function
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//(d, deck) is receiver, Any variable of type deck, now gets access to print function with this receiver function
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
