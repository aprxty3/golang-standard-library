package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"John", "Jane", "Doe", "Smith"}
	highscore := []int{10, 20, 30, 40}

	fmt.Println(slices.Max(names))
	fmt.Println(slices.Max(highscore))
	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(highscore))
	fmt.Println(slices.Contains(names, "Doe"))
	fmt.Println(slices.Contains(highscore, 30))
	fmt.Println(slices.Index(names, "Doe"))
	fmt.Println(slices.Index(highscore, 30))
}
