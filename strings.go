package main

import (
	"fmt"
	"strings"
)

func main()  {
	fmt.Println(strings.Contains("Hello, World!", "World"))
	fmt.Println(strings.Count("Hello, World!", "o"))
	fmt.Println(strings.Replace("Hello, World!", "World", "Go", 1))
	fmt.Println(strings.Split("Hello, World!", ","))
	fmt.Println(strings.Join([]string{"Hello", "World"}, ","))
	fmt.Println(strings.ToUpper("Hello, World!"))
	fmt.Println(strings.TrimSpace(" Hello, World! "))
	fmt.Println(strings.HasPrefix("Hello, World!", "Hello"))
	fmt.Println(strings.HasSuffix("Hello, World!", "World"))
	fmt.Println(strings.Index("Hello, World!", "World"))
	fmt.Println(strings.LastIndex("Hello, World!", "o"))
	fmt.Println(strings.Repeat("Hello, World!", 3))
	
}