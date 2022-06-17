package main

import "fmt"







func main() {

	customer := getData(1)
	fmt.Println(customer) 


} 





func getData() (customers []string ) {

	// create 1 record 

	customers := []string
	customers = append(customers, "Jahid Ahmed")

}








// This function return some data, it could be coming from database/file.
// Function can take multiple inputs and return multiple output

// func getData(input)(output){
//	}


// Control Dlows (if/ese)
// Control flows allow us to add "rules" to our code
// if x is this, do that, else do something else


// Arrays 

// Arrays allow us to make a collection of variable of the same type 
// Downsize in arrays, its fixed in size 
 
 // Slices 

 //Slices are dynamically-sized view into arrays. 


//func getData(customerId int) (customers [2]string ) {

//	 firstname := "Jahid"0
//	 lastname  := "Ahmed"
//	 return firstname + " " + lastname
//   return "Jahidoo"
//
//	if customerId == 1 {
//		return "Jahid Ahmed"
//	}   else if customerId  == 2 {
//		return "Marcel Dempers"		
//	}	else {
//		return ""
//	} 

//}




//func getData() (customers [2]string ) {

	// create 1 record 

//	customer := "Jahid Ahmed"
//	customers[0] = customer
//	customers[1]
//	return customers
//}


// To Do 

// - Loops