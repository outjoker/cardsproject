package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

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

func (d deck) toString() string {
	// this functions takes deck (slice of strings) as input and convert that as a string and return it
	// converting deck to slice of strings

	return strings.Join([]string(d), ",") // this will join the strings in the slice with comma as the separator
}

//save the decks to local file system
// 1. take the dec and convert it to a slice of strings
// 2. join the slice of strings and convert it to a string with commma as the separator
// 3. convert the string to byte slice []byte(string)
// we can call the function in this manner, cards.saveToFile, if we implement this function as receiver function
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666) // the last param is file permission currently it is 0666 which is all can read and write

}
