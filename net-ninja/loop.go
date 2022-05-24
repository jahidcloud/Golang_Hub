package main

import (
	"fmt"
)

func main() {

	i := 1

	for i <= 5 {
		fmt.Println("Hello World", i)
		i++
		//i = i+1

	}

	for x := 1; x <= 3; x++ {
		fmt.Println("Hello London")
	}

	names := []string{"mari", "luigi", "yoshi", "peach"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

}
