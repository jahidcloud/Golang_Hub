package main 


import (
    "errors"
    "fmt"
)    

func main() {
    err := returnError()
    fmt.Println(err)
    //     if err != nil {
//        fmt.Println(err)
 //    }

}



func returnError(returnError bool)  (string,error) {
    return errors.New("Error alert")
}
