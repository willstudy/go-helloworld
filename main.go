package main

import (
	"fmt"
	"net/http"
)

func Hello(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		fmt.Fprintf(response, "Hello World by GET method\n")
		break
	case "POST":
		fmt.Fprintf(response, "Hello World by POST method\n")
		break
	case "PUT":
		fmt.Fprintf(response, "Hello World by PUT method\n")
		break
	case "DELETE":
		fmt.Fprintf(response, "Hello World by DELETE method\n")
		break
	default:
		fmt.Fprintf(response, `{"error":1, "message":"unsupport method"}`)
		break

	}
}

func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":9000", nil)
}
