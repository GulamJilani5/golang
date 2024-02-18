package main

import "fmt"

func main() {

	cards := []string{"Ace of Diamond", newCard()}

	cards = append(cards, "Six of Spades")

	// fmt.Println(cards)

	for i, card := range cards {
		fmt.Println(card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
