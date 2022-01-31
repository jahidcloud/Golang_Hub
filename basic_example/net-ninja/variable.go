package main 

import (

	"fmt"
)



func main(){
	// strings

	var name string = "eren"
	var char_age int 	= 21
	pet := "Dog"
	fmt.Println(name, char_age)
	fmt.Println(char_age)
	fmt.Println(pet)

	// ints 

	var age int = 22
	var day = 15 
	year := 2004

	fmt.Println(age, day, year)


	// bits & memory 

	var numOne int8 = 15 // int8 range: -128 - 127
	var numTwo int8 = -128
	var numThree uint = 25 // uint means you cannot assign a negative valie
	// Most of the time, you do not have to specify 

	fmt.Println(numOne, numTwo, numThree)


	// Float 
	// They are decimal. Higher amount of Bits, larger the range.

	var scoreOne float32 = 2.5
	fmt.Println(scoreOne)


}


