package main

import "fmt"

func main() {
	cards := deck{"Ace of Diamonds", newCards()}
	cards = append(cards, "Six of Spades")
	fmt.Println(cards)

	cards.print()
}

func newCards() string {
	return "Five of Diamonds"
}
