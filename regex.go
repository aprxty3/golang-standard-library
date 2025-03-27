package main

import (
	"fmt"
	"regexp"
)

func main() {
	regex := regexp.MustCompile(`[a-z]+`)

	fmt.Println(regex.MatchString("halo"))
	fmt.Println(regex.MatchString("123"))

	fmt.Println(regex.FindAllString("halo 123", -1))
}
