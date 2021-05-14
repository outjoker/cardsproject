package main

import (
	"fmt"
	"testing"
)

func TestNewDeck(t *testing.T) {
	// t in the args section is the test handler
	// create new deck
	// create a conditional block to check if the deck has correct number of cards
	// in else case we need to say the go test handler that something went wrong

	d := newDeck()

	fmt.Println(len(d))
	if len(d) != 16 {
		t.Errorf("expected length of the deck as 17, but received %v", len(d))
	}
}

func TestFirstCardInTheDeckIsAceOfSpaces(t *testing.T) {

	d := newDeck()
	if d[0] != "Ace of Spades" {
		t.Errorf("expected first card to be Ace of Spades, but got %v", d[0])
	}

}

func TestLastCardInTheDeckIsFourOfClubs(t *testing.T) {

	d := newDeck()
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("expected last card to be Four of Clubs, but got %v", d[len(d)-1])
	}

}
