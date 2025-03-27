package main

import (
	"fmt"
	"time"
)

func main() {
	duration1 := 100 * time.Second
	duration2 := 10 * time.Millisecond
	duration3 := 1 * time.Hour

	fmt.Println(duration1)
	fmt.Println(duration2)
	fmt.Println(duration3)

	fmt.Printf("Duration1: %d\n", duration1)
	fmt.Printf("Duration2: %d\n", duration2)
	fmt.Printf("Duration3: %d\n", duration3)
}
