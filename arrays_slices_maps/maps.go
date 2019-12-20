package main

import "fmt"

func main() {
	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x["key"])
	delete(x, "key")
	fmt.Println(x["key"])
	dict2 := make(map[int]int)
	dict2[1] = 12
	fmt.Println(dict2[1])
}
