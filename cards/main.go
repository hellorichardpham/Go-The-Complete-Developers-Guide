package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	// variable reassignment should be = and not :=. := is only for inferred initialization
	card := "Ace of Spades"
	card = "Five of Diamonds"
	fmt.Println(card)
}
