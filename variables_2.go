package main

import "fmt"

func main() {
	var x string = "Hello"
	var y string = "world"
	fmt.Println(x == y)
	y = "Hello"
	fmt.Println(x == y)
	z := 5
	fmt.Println(z)
}
