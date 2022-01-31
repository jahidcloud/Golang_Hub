// 	Post Function takes three parametes;
// 		- URL address of the server 
//		- The content type of the body as string 
//		- The request body that is to be sent using POST method of type io.Reader

// Post request to (https://postman-echo.com/post)




postBody, _ := json.Marshal(map[string]string{
	"name": "Tobdy",
	"email": "Toby@example.com",

})

responseBody := bytes.newBuffer(postBody)