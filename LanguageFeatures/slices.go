package main

import "fmt"

func main()  {
  // a slice is a sort of array that can expand or contract
  movies  := []string{"Dr.No", bondMovie()}
  movies = append(movies, "Moonraker")
  for i, movie  := range movies {
  fmt.Println(i, "-", movie)
  }
}

func bondMovie() string {
	return "You only live twice"
}
