package main

func main() {
	cards := newDeck()

	cards.print()

	hand, remaindingCards := deal(cards, 4)

	hand.print()
	remaindingCards.print()
}
