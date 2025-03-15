package main

import (
	"flag"
	"fmt"
)

func main() {
	username := flag.String("username", "default_user", "username")
	age := flag.Int("age", 20, "age")

	flag.Parse()

	fmt.Println("Username:", *username)
	fmt.Println("Age:", *age)
}
