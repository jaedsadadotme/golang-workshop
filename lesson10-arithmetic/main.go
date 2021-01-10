package main

import (
	"fmt"
	"strconv"
)

func convert() {
	celsius := 35
	fahrenheit := (9*celsius + 160) / 5

	fmt.Printf("celsius: %d is %d fahreheit ", celsius, fahrenheit)
}

func incdecStatement() {
	n := 10
	fmt.Println(n)
	n += 1
	fmt.Println(n)
	n -= 1
	fmt.Println(n)
	n++
	fmt.Println(n)
}
func stringConvert() {
	var a string
	feet, _ := strconv.ParseFloat(a, 64)
	fmt.Println(feet)
}

func main() {
}
