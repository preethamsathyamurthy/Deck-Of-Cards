package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type deck
// which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	suits := []string{"Spades", "Diamonds", "Cloves", "Hearts"}
	values := []string{"Ace", "Two", "Three",
		"Four", "Five", "Six", "Seven",
		"Eight", "Nine", "Ten",
		"Jack", "Queen", "King",
	}

	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	cardBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option 1 - Log the error and return a call to newDeck()
		// Option 2 - Log the error and entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	cardString := string(cardBytes)
	return deck(strings.Split(cardString, ","))
}

func (d deck) shuffle() {

	// make things truly random by
	// generating a random number with seed
	// of current timestamp
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		// swap numbers based on random number generator
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
