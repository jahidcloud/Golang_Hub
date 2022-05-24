package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)
	fmt.Println(m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

}
