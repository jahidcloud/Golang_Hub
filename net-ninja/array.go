package main

import "fmt"

func main() {
	var ages [3]int = [3]int{20, 25, 30} // fixed length in array in go
	names := [4]string{"Yoshi", "mario", "peach", "bowser"}
	names[1] = "luigi"

	fmt.Print(ages)
	fmt.Println(names, len(names))
}
