package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	//Test the length of deck of cards
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	//Test the first card of the deck
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades, but received %v", d[0])
	}
	//Test the last card of the deck
	if d[len(d)-1] != "Jack of Clubs" {
		t.Errorf("Expected ack of Clubs, but received %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	//Clean any previously created file
	os.Remove("_decktesting")
	//Create new Deck
	d := newDeck()
	//Save deck to file
	d.saveToFile("_decktesting")
	//Load new Deck from this saved file
	newLoadedDeck := newDeckFromFile("_decktesting")
	//Check the length of new loaded deck
	if len(newLoadedDeck) != 52 {
		t.Errorf("Expected 52 cards, got %v", len(newLoadedDeck))
	}
	//Remove any decktesting file if left
	os.Remove("_decktesting")
}
