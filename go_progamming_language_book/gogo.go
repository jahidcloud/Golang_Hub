package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "hi, this is, educative" // removes the comma
	split := strings.Split(str, ",")
	fmt.Println(split)
	fmt.Println("The length of the slice is:", len(split))

}

//The Split() method​ in Golang (defined in the strings library)
// breaks a string down into a list of substrings using a specified separator. The method returns the substrings in the form of a slice.

// Using split method, we have to give two arguement which is the string values and the delimiter(seperator)
// str is the string to be split.
