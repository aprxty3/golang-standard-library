package main

import (
	"errors"
	"fmt"
)

var (
	ValidationErr = errors.New("validation error")
	NotFoundErr   = errors.New("not found")
)

func GetId(id string) error {
	if id == "" {
		return ValidationErr
	}
	if id != "123" {
		return NotFoundErr
	}
	return nil
}

func main() {
	err := GetId("")
	if err != nil {
		if errors.Is(err, ValidationErr) {
			fmt.Println("Validation error")
		} else if errors.Is(err, NotFoundErr) {
			fmt.Println("Not found error")
		} else {
			fmt.Println("Unknown error")
		}
	} else {
		fmt.Println("Success")
	}
}
