package main 

import (
	"os"
	"fmt"
)


func main(){
	file, err := os.Open("file.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("the file is: ", file) // prints out pointer 
}

