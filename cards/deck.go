package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

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

//Save Deck to a file
func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

//(d, deck) is receiver, Any variable of type deck, now gets access to print function with this receiver function
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//Convert the deck to string, receiver is better here
func (d deck) toString() string {
	//Step 1: convert deck to slice of string ([]string(d))
	//Step 2: convert slice of string to a string, values seperated by comma
	//["red","yellow","pink"] --> ["red,yellow,pink"]
	return strings.Join([]string(d), ",")
}

//Create a new Deck from a file from a hdd
func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		//Option -1: log the error and return a new deck OR
		//Option -2: log the error and entirely quit the program
		//We will implement option -2
		fmt.Println("Error:", err)
		//call the os exit function
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), ","))
}
