package main

import "fmt"

func main() {
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["Ne"] = "Neon"
	elements["O"] = "Oxygen"
	fmt.Println(elements["Ne"])
	fmt.Println(elements["Li"])
}
