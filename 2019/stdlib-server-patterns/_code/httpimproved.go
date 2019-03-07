// +build OMIT
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// START_INFO OMIT

// Info is just a logger with extra steps
var Info *log.Logger = log.New(os.Stdout,
	"INFO: ",
	log.Ldate|log.Ltime)

//STOP_INFO OMIT

// START_STRUCT OMIT

// Response is a model to server responses
type Response struct {
	Message string
	Code    int
}

// STOP_STRUCT OMIT

// START_HOMEHANDLER OMIT

// Home is the handler of hello world, the web server // OMIT
func Home(w http.ResponseWriter, req *http.Request) {

	if req.URL.Path != "/" { // Reject requests to different paths
		http.Error(w, "404 not found.", http.StatusNotFound) // OMIT
		return                                               // OMIT
	} //OMIT

	switch req.Method {
	case "GET":
		data := []Response{Response{"Hello GET", 200}}
		b, _ := json.Marshal(data)

		Info.Println("Sending response", string(b)) // OMIT
		io.WriteString(w, string(b))
	case "POST":
		fmt.Fprintf(os.Stdout, "Post from website! req.PostFrom = %v\n", req.PostForm)
		Info.Println("Post from website from", req.PostForm) // OMIT
		data := []Response{Response{"Hello POST", 200}}
		b, _ := json.Marshal(data)

		Info.Println("Sending response", string(b)) // OMIT
		io.WriteString(w, string(b))
	default:
		Info.Println("Sorry, only GET and POST methods are supported.") // OMIT
		io.WriteString(w, "Sorry, only GET and POST methods are supported.")
	}
}

// STOP_HOMEHANDLER OMIT

// START_MAIN OMIT
func main() {
	http.HandleFunc("/", Home)

	Info.Println("Listening at 5555")

	log.Fatalf("INFO: ", log.Ldate|log.Ltime, http.ListenAndServe(":5555", nil))
}

// STOP_MAIN OMIT
