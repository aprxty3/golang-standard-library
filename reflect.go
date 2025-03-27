package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

type Person struct {
	Name    string `required:"true" max:"10"`
	Address string `required:"true" max:"10"`
	Age     int    `required:"true" max:"10"`
}

func readField(value any) {
	valueType := reflect.TypeOf(value)
	fmt.Println("Type Name: ", valueType.Name())

	for i := 0; i < valueType.NumField(); i++ {
		valueField := valueType.Field(i)
		fmt.Println(valueField.Name, "with type", valueField.Type)
		fmt.Println("Required: ", valueField.Tag.Get("required"))
		fmt.Println("Max: ", valueField.Tag.Get("max"))
	}

}

func isValid(value any) (result bool) {
	result = true
	t := reflect.TypeOf(value)
	v := reflect.ValueOf(value)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			fieldValue := v.Field(i).Interface()
			result = fieldValue != ""
			if result == false {
				return result
			}
		}
	}
	return result
}

func main() {
	readField(Sample{Name: "John Doe"})
	readField(Person{Name: "Jane Doe", Address: "Jakarta", Age: 30})

	person := Person{Name: "Aji", Address: "Jakarta", Age: 30}
	fmt.Println("Valid: ", isValid(person))

	person1 := Person{Name: "Rhue", Address: "", Age: 3}
	fmt.Println("Valid: ", isValid(person1))
}
