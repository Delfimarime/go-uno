package main

import (
	"../internal"
	"fmt"
)

func main() {
	deck := internal.Deck{}
	deck.AfterPropertiesSet()
	fmt.Println(deck)
	fmt.Print(deck.Size())
}
