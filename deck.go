package main

import "fmt"

// create a type
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuites := []string{"Spades", "Diamond", "Heart", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards

}

// this is receiver method of type deck
func (d deck) print() {
	for _, value := range d {
		fmt.Println(value)
	}
}

// this is independent function
func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}
