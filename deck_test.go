package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected a deck of 52 cards, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected the first card to be Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "King of Hearts" {
		t.Errorf("Expected the last card to be King of Hearts but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_testingDeck.txt")
	deck := newDeck()

	deck.saveToFile("_testingDeck.txt")

	loadedDeck := newDeckFromFile("_testingDeck.txt")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected a deck of 52 cards, but got %v", len(loadedDeck))
	}

	if loadedDeck[0] != "Ace of Spades" {
		t.Errorf("Expected the first card to be Ace of Spades, but got %v", loadedDeck[0])
	}

	if loadedDeck[len(loadedDeck)-1] != "King of Hearts" {
		t.Errorf("Expected the last card to be King of Hearts but got %v", loadedDeck[len(loadedDeck)-1])
	}

	os.Remove("_testingDeck.txt")
}
