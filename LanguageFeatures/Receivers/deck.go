package main

import "fmt"

// create a  new type called deck, using the slices we learned

type deck []string

func newDeck() deck {
	cards := deck{}

	// cards are a combination of values in a given suit - which is really a bad English transliteration of 'suite', because they 'follow' each other, but whatever.
	cardSuits := []string { "Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string {"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Knave", "Lady", "King"}

	// mix and match
	for _, suit := range  cardSuits {
		for  _, value := range cardValues {
			cards = append(cards, suit + " of " + value)
		}
	}
	return cards
}

// functions know what they work on based on the parameter defined after the parenthesis
func(d deck) print() {
	// by convention, go-fluent writers define the reference for the type of the receiver function with a one or two letter word (here, d, like deck)
	for i,c := range d {
	 fmt.Println(i, c)
	}
}
