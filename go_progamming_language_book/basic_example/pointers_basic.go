package main 


import (
	"fmt"
)




func update_name (x *string){
	*x = "Eren"
//	fmt.Println(n)
//	p := &n
//	fmt.Println(p)


}



func main(){
	n := "armin"
	
//	update_name(n)

//	fmt.Println("memory address of 'n' is: ", &n)

	m := &n
//	fmt.Println("Memory address:", m)
//	fmt.Println("value at memory address:", *m)

//	fmt.Println(n)
//	fmt.Println(*m)

	fmt.Println(n)
	update_name(m)
	fmt.Println(n)


}

//https://www.youtube.com/watch?v=4B2rwYvuiBo&list=PL4cUxeGkcC9gC88BEo9czgyS72A3doDeM&index=14