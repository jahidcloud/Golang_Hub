package main

import (
	"fmt"
	"os"
)

func main() {

	// The first argument
	// is always program name
	myProgramName := os.Args[1]
	myProgramName2 := os.Args[0]
	gettingArgs := os.Args[2]

	// it will display
	// the program name
	fmt.Println(myProgramName)
	fmt.Println(myProgramName2)
	fmt.Println(gettingArgs)

}
