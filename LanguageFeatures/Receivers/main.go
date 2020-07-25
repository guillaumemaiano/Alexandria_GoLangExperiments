package main

import "fmt"

func main() {
	cards := deck {
	"one",
	"two",
	}
	/* The following one-liner is equivalent to:
	for i, card := range cards {
		fmt.Println(i, card)
	}
	*/
	cards.print()
}
