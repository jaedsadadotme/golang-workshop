package main

import (
	"fmt"
	"unicode/utf8"
)

func demostring() {
	var a string
	a = "hello" // string single line
	a = `hello` // raw string can multi line
	fmt.Println(a)
}
func countString() {
	name := "puen"
	fmt.Println(len(name)) // normal count
	name = "ปืน"
	fmt.Println(utf8.RuneCountInString(name)) // count with utf8
}
func main() {
	countString()
}
