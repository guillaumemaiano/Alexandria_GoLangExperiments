package main

// create a  new type called deck, using the slices we learned

type deck []string

// functions know what they work on based on the parameter defined after the parenthesis
func(d deck) print() {
	for i,c := range d {
	 fmt.Println(i, c)
	}
}
