package main

import "fmt"

func name(x string) {

	fmt.Println(x)
}

func main() {
	x := 7
	y := &x
	fmt.Println(x, y)
	*y = 8
	fmt.Println(x, y)

}
