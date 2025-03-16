package main

import (
	"fmt"
	"strconv"
)

func main() {
	result, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("Error parsing bool:", err)
	} else {
		fmt.Println("Result:", result)
	}

	num, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("Error parsing int:", err)
	} else {
		fmt.Println("Result:", num)
	}

	num2 := strconv.Itoa(123)
	fmt.Println("Result:", num2)

	num3, err := strconv.ParseFloat("123.45", 64)
	if err != nil {
		fmt.Println("Error parsing float:", err)
	} else {
		fmt.Println("Result:", num3)
	}

	num4 := strconv.FormatInt(123, 10)
	fmt.Println("Result:", num4)

	num5 := strconv.FormatFloat(123.45, 'f', -1, 64)
	fmt.Println("Result:", num5)
	
	
}
