package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (s UserSlice) Len() int {
	return len(s)
}

func (s UserSlice) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func (s UserSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	users := []User{
		{"John", 30},
		{"Jane", 20},
		{"Doe", 40},
		{"Smith", 50},
		{"Doe", 10},
		{"Smith", 60},
	}

	sort.Sort(UserSlice(users))
	fmt.Println(users)

}
