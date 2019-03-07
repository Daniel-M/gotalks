// +build OMIT
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"encoding/json"

	"github.com/gorilla/mux"
)

// START_CONST OMIT
const port = ":8100"

// STOP_CONST OMIT

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

// START_HOMEGETHANDLER OMIT
// HomeGET handles GET requests to the hello world, the web server
func HomeGET(w http.ResponseWriter, req *http.Request) {

	if req.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	data := []Response{Response{"Hello GET", 200}}
	b, _ := json.Marshal(data)

	Info.Println("Sending response", string(b))
	io.WriteString(w, string(b))
}

// STOP_HOMEGETHANDLER OMIT

// START_HOMEPOSTHANDLER OMIT

// HomePOST handles POST requests for hello world, the web server
func HomePOST(w http.ResponseWriter, req *http.Request) {

	if req.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(os.Stdout, "Post from website! req.PostFrom = %v\n", req.PostForm)
	Info.Println("Post from website from", req.PostForm)
	data := []Response{Response{"Hello POST", 200}}
	b, _ := json.Marshal(data)

	Info.Println("Sending response", string(b))
	io.WriteString(w, string(b))
}

// STOP_HOMEPOSTHANDLER OMIT

// START_MAIN OMIT
func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", HomeGET).Methods("GET")
	r.HandleFunc("/", HomePOST).Methods("POST")

	fmt.Println("The router is running at", port)

	http.Handle("/", r)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "An Error has ocurred: %s\n", err)
	}
}

// STOP_MAIN OMIT
