package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

func (d deck) print() {
	// this is receiver method of type deck for printing deck
	for _, value := range d {
		fmt.Println(value)
	}
}

func (d deck) toString() string {
	// this is receiver method of type deck for converting deck to string
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	// this is receiver method of type deck for saving deck to file
	data := []byte(d.toString())
	return os.WriteFile(filename, data, 0644)
}

func newDeckfromFile(filename string) deck {
	// this is independent function for loading deck from file
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// type conversion and split
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func deal(d deck, handsize int) (deck, deck) {
	// this is independent function
	return d[:handsize], d[handsize:]
}

func (d deck) shuffle() {
	// create source object and provide seed value (unixNano is int representation of current time)
	source := rand.NewSource(time.Now().UnixNano())
	// create object of type Rand
	r := rand.New(source)

	for indx, _ := range d {
		rand_num := r.Intn(len(d) - 1)
		d[indx], d[rand_num] = d[rand_num], d[indx]
	}
}
