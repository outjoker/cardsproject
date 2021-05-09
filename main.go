package main

import "fmt"

var deckSize int

func main() {

	deckSize = 52
	fmt.Println(deckSize)

	//cards := []string{"joker", "ace", "spade"}
	// since in deck.go we have declared a custom type "deck" which is an
	// slice of strings, i can change the cards definition above to use
	// the type deck instead of using []string

	// cards := deck{"joker", "ace", "spade"}
	// cards = append(cards, "six of spades")

	// for i, each_card := range cards {
	// 	fmt.Println(i, each_card)
	// }

	cards := newDeck()
	//cards.print()

	hand, remainingCardsInDeck := deal(cards, 5)
	fmt.Println("-------------------------------------------------printing the hand")
	hand.print()
	fmt.Println("------------------------")

	fmt.Println("-------------------------------------------------printing the remaining cards in the deck")
	remainingCardsInDeck.print()
	fmt.Println("------------------------")

}

func newCard() string {
	return "5 of diamonds"
}
