package main

func main() {
	// cards := newDeck()

	// hand, remainingInDeck := deal(cards, 10)

	// fmt.Println("Hand")
	// hand.print()
	// fmt.Println("Remaining In Deck")
	// remainingInDeck.print()

	// fmt.Println(cards.toString())

	// cards.saveToFile("firstfile.txt")

	cards := newDeck()
	cards.shuffle()
	cards.saveToFile("firstfile.txt")

	cards_loaded := newDeckFromFile("firstfile.txt")
	cards_loaded.print()
}
