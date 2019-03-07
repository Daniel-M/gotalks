// START_HOMEHANDLERREQ OMIT

// Home is the handler of hello world, the web server
func Home(w http.ResponseWriter, req *http.Request) {

	if req.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	// STOP_HOMEHANDLERREQ OMIT

	// START_REQMETHOD OMIT

	switch req.Method {
	case "GET":
		// Code for deal with GET

	case "POST":
		// Code for deal with POST 

	default:
		// Fallback responses
	}
	// STOP_REQMETHOD OMIT

}

