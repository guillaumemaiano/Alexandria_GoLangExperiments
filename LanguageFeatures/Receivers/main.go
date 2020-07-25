package main


func main() {
	cards := deck {
	"one",
	"two",
	}
	/* The following one-liner is equivalent to:
	// import fmt
	for i, card := range cards {
		fmt.Println(i, card)
	}
	*/
	cards.print()
}
