package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	// Port TLS should be served from
	TLS_PORT = ":443"

	// Port HTTP should be served from
	HTTP_PORT = ":8077"

	// Information pertaining to the nature in which the server's log file
	// may be opened
	LOG_FILE_FLAGS = os.O_RDWR | os.O_CREATE | os.O_APPEND
	LOG_FILE_MODE  = 0666

	DOMAIN_NAME = "localhost"

	// Location of HTML pages
	PAGES_DIR = "../pages"
)

var (
	LOG_FILE_PATH = os.Getenv("LOG_FILE_PATH")

	// Get certificate information from environment variables
	CERT_KEY_PATH   = os.Getenv("CERTIFICATE_KEY_PATH")
	CERT_CHAIN_PATH = os.Getenv("CERTIFICATE_CHAIN_PATH")

	ENDPOINT_HANDLER_MAP = map[string]func(http.ResponseWriter, *http.Request){
		"/":                indexHandler,
		"/" + "gpg":        gpgHandler,
		"/" + "blog" + "/": blogHandler,
	}
)

// pageNotFoundHandler is used to display a 404 erro page when a visitor tries
// to view a page that does not exist.
func pageNotFoundHandler(writer http.ResponseWriter, request *http.Request) {
	http.ServeFile(writer, request, PAGES_DIR+"/errors/404.html")
}

func blogHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello from the blog handler!")
}

// indexHandler is a catch all handler for url's that do not match any other
// handler. Will display the index page.
func indexHandler(writer http.ResponseWriter, request *http.Request) {
	// That page doesn't exist
	if request.URL.Path != "/" {
		pageNotFoundHandler(writer, request)
		return
	}

	http.ServeFile(writer, request, PAGES_DIR+"/index.html")
}

// gpgHandler handles requests for gpg.skippola.com. Will return my public GPG
// key.
func gpgHandler(writer http.ResponseWriter, request *http.Request) {
	http.ServeFile(writer, request,
		PAGES_DIR+request.URL.Path+"/kai_fleischman.gpg")
}

// redirectHttpToHttps redirects all incoming HTTP requests to the HTTPS
// equivalent.
func redirectHttpToHttps(writer http.ResponseWriter, request *http.Request) {
	// remove/add not default ports from req.Host
	target := "https://" + request.Host + request.URL.Path
	if len(request.URL.RawQuery) > 0 {
		target += "?" + request.URL.RawQuery
	}

	log.Printf("redirect to: %s", target)
	http.Redirect(writer, request, target, http.StatusTemporaryRedirect)
}

func main() {
	// Setup the logger to write logs to a file
	log_file, err := os.OpenFile(LOG_FILE_PATH, LOG_FILE_FLAGS, LOG_FILE_MODE)

	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}

	// Defer closing the file until the main function is done
	defer log_file.Close()

	// Set Go's logger to write logs to the server's log file
	log.SetOutput(log_file)

	// Setup handlers for certain endpoints using the ENDPOINT_HANDLER_MAP
	for endpoint, handler := range ENDPOINT_HANDLER_MAP {
		http.HandleFunc(endpoint, handler)
	}

	// Redirect all HTTP traffic to HTTPS
	err = http.ListenAndServe(HTTP_PORT, nil) //http.HandlerFunc(redirectHttpToHttps))

	log.Printf("Handling traffic for skippola.com on port%s", TLS_PORT)
	//err = http.ListenAndServeTLS(TLS_PORT, CERT_CHAIN_PATH, CERT_KEY_PATH, nil)
	log.Fatal(err)
}
