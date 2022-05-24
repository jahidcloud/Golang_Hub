package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Enter a single Character: ")
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(char)

}
