package main 


import (
	"fmt"
	"log"
	"net/http"
)


funchelloHandler(w http.ResponseWriter, r "http.Request"){
	
}


func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHanlder)\


	fmt.Printf("Starting server at port 8-8-\n")

	if err := http.ListenAndServe(":8080", nil); err !=nil {
		log.Fatal(err)
	}
	 

}