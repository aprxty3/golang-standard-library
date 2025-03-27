package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println(currentTime)

	utc := time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println(utc)
	fmt.Println(utc.Local())

	formatter := "2006-01-02 15:04:05"
	value := "2020-01-01 00:00:00"
	// value := "ASAL"
	timeParsed, err := time.Parse(formatter, value)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(timeParsed)
	}

	fmt.Println(timeParsed.Year())
	fmt.Println(timeParsed.Month())
	fmt.Println(timeParsed.Day())
	fmt.Println(timeParsed.Hour())

}
