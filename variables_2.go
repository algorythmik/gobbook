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
	// constants
	const constantString string = "Can not be changed"
	fmt.Println(constantString)
	// declaring multiple variables
	var (
		a = 10
		b = 20
		c = 30
	)
	fmt.Println(a + b + c)

}
