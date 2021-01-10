package main

import "fmt"

func main() {
	var name string = "asdfsda"

	a := 12312            // short declare and type int
	a, b := 12312, "asdf" // mutiple short declare
	fmt.Println(name, a, b)
}
