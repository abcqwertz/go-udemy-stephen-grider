package main

import "fmt"

func main() {
	cards := newDeck()
	hand, remainingCars := deal(cards, 5)
	fmt.Println(hand, remainingCars)
}
