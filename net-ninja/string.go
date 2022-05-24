package main

import (
	"fmt"
	"strings"
)

func main() {
	greeting := "hello there friends !"
	fmt.Print(greeting)

	fmt.Println(strings.Contains(greeting, "hello"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))

}
