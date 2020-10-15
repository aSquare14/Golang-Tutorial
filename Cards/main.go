package main

import "fmt"

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.saveToFile("my_cards")
	fmt.Println(cards.toString())
}
