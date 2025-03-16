package main

import (
	"container/list"
	"fmt"
)

func main()  {
	data := list.New()

	data.PushBack("Aji")
	data.PushBack("Prasetyo")
	data.PushBack("Saputra")

	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	data.Remove(data.Front())

	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}
	
	
}