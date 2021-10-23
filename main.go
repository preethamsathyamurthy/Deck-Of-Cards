package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"

	// card := "Ace of Spades"
	// card = "Five of Diamonds"

	card := newCard()
	fmt.Println(card)
}

// Files in the same package can freely call functions defined in other files.
// go run main.go newcard.go
