package main

import (
   "io/ioutil"
   "log"
   "net/http"
)


// Basic Example of Get Request

func main(){
	 resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
 	if err != nil {
		 log.Fatalln(err)
	 }

	//We Read the response body on the line below.
	 body, err := ioutil.ReadAll(resp.Body)
	 if err != nil{
		 log.Fatalln(err)
	 }

	 //Convert the body to type string
	 sb := string(body)
	 log.Printf(sb)	
}



