package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

// after getting the filename as input,this function should return a deck accordingly
func getNewDeckFromFile(fileName string) deck {
	byteSliceReceived, error := ioutil.ReadFile(fileName)
	// byteSliceReceived will be the string of chars/ string of cards separated by comma
	// error will catch the value of type error, if all is well error will be nil
	if error != nil {
		fmt.Println("error in reading the file: ", error)
		// if we are not able to load a deck from file, we can call newDeck() declared on top as a fallback approach
		os.Exit(1) // reason 1 is being passed is, if Exit(0)=> it indicates no error and vice-versa
	}
	// convert the byteslice => string => slice of strings => deck
	// string(byteSliceReceived) // the value would be Ace of spaces, two of diamonds
	// now we should convert this whole string to slice of strings, so we use strings.Split(string, delimiter)
	sliceOfStrings := strings.Split(string(byteSliceReceived), ",")
	return deck(sliceOfStrings)
}

func (d deck) shuffleCards() {
	// this function receives a deck as it is a receiver

	// generating random source which generates the random numbers, by making the source random, everytime the generation also becomes
	// random
	// this source shouldn't be the same whenever we run this program
	// for that matter we can use the UnixNano() of the time package
	// which means everytime we run this program, it will take the current time*/
	source := rand.NewSource(time.Now().UnixNano())
	randomNumber := rand.New(source) // the value of randomNumber is of type rand and it can use rand functions

	for indexOfCard := range d {
		//randomPosition := rand.Intn(len(d) - 1)
		randomPosition := randomNumber.Intn(len(d) - 1)

		// swapping card at index as the randomposition with card at current index in the slice of decks
		d[indexOfCard], d[randomPosition] = d[randomPosition], d[indexOfCard]

	}

}
