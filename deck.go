package main

import "fmt"

// create a new custom type of `deck` which is slice of strings
type deck []string

func newDeck() deck {
	// this function should return a list of cards in the deck
	// why did we annotate deck in function declaration? because we are returning slice of strings / slice of cards
	cards_in_deck := deck{}

	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}

	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, each_suite := range cardSuites {
		for _, each_card_value := range cardValues {
			cards_in_deck = append(cards_in_deck, each_card_value+" of "+each_suite)
		}
	}

	return cards_in_deck
}

func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}

// this function is supposed to be called with existing list of cards
// so that the functions returns a hand based on the handsize provided in the params
func deal(d deck, handSize int) (deck, deck) {

	return d[:handSize], d[handSize:]

}
