package main


import (
    "fmt"
    "log"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/hello" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }

    if r.Method != "GET" {
        http.Error(w, "Method is not supported.", http.StatusNotFound)
        return
    }


    fmt.Fprintf(w, "Hello!")
}


func formHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil {
		fmt.Fprint(w, "ParseForm () err: %v", err)
	}

	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprint(w, "Name = %s\n", name)
	fmt.Fprint(w, "Address = %s\n", address)

}




func main() {
	fileServer := http.FileServer(http.Dir("./Static"))
	http.HandleFunc("/form", formHandler)
	http.Handle("/", fileServer)
    http.HandleFunc("/hello", helloHandler) // Update this line of code


    fmt.Printf("Starting server at port 8080\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}