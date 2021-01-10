package main

import "fmt"

var x int  // declare varieble global and unexport
var Xx int // declare varieble global and export
func main() {
	var y int
	_ = y // declare _ for varieble is not use. it's like black hole

	var (
		z   int
		zz  bool
		zzz string
	) // declare mutiple varieble
	fmt.Println(z, zz, zzz)

	var a, aa int // this is parallel mutiple declaration
	fmt.Println(a, aa)
}
