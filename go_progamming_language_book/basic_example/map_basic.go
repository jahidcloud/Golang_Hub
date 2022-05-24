package main 


import (
	"fmt"
)

// maps are made of key/values pairs.

func main (){

menu := map[string]float64{
	"soup":  			4.99,
	"pie": 	 			7.99,
	"salad": 			6.99,
	"toffee pudding": 	3.55,

}

fmt.Println(menu)
fmt.Println(menu["pie"]) // we have to pass in strings in quotes 

// looping maps 

for k, v := range menu {
	fmt.Println(k, "-", v)
	
}


	// int as key type 

	phonebook := map[int]string {
		7677224133:	"mario",
		9847593733:	"luigi",
		8457754854: "peach",
	}

	fmt.Println("\n", phonebook)
	fmt.Println("\n", phonebook[7677224133])

}