package main

// to look through Go packages: https://golang.org/pkg
// for example -> import ioutil // disk operations

func main() {
	cards := newDeck()
	
	/* The following one-liner is equivalent to:
	// import fmt
	for i, card := range cards {
		fmt.Println(i, card)
	}
	*/
	cards.print()
	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()
}
