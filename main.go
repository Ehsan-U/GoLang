package main

func main() {
	cards := newDeckfromFile("deck")
	// cards.print()
	cards.shuffle()
	cards.print()

}
