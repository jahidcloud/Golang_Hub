package main 


import (
	"fmt"
	"log"
	"net/http"
	"os"
)



func main() {
	url := os.Args[1]

	api(url)


}


//Passing an Parameter called URL §
func api(url string){
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response)  // if no error, this response will be printed

}


// Gonna allow us to pass OS arguement and call the api function