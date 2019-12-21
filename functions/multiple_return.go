package main

import "fmt"

func main() {
	x, y := f(2)
	fmt.Println(x, y)
}

func f(x int) (int, int) {
	return x + 2, x + 3

}
