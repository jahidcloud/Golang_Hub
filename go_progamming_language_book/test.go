package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	argone := os.Args[1]
	arg2 := os.Args[2]
	fmt.Println("Hello:", argone, arg2)

	all_arg := argone + arg2
	words := strings.Fields(all_arg)

	fmt.Println(words)
}

//f, err := os.Open(name)
//if err != nil {
//return err
//}
// ...use f...
//f.Close()
