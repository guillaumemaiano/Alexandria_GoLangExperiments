package main

import "fmt"

func main() {
	cards := deck {
	"one",
	"two",
	}

	for i, card := range cards {
		fmt.Println(i, card)
	}
}
