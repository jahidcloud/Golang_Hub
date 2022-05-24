package main

import "fmt"

func sumAB(a int, b int) int {
	var result int
	result = a + b
	return result

}

func main() {
	var result int = sumAB(9, 10)
	fmt.Println(result)
}
