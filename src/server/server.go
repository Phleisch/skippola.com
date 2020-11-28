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

	// Information pertaining to the nature in which the server's log file
	// may be opened
	LOG_FILE_FLAGS = os.O_RDWR | os.O_CREATE | os.O_APPEND
	LOG_FILE_MODE  = 0666

	DOMAIN_NAME = "localhost"
	PAGES_DIR = "../pages"
)

var (
	LOG_FILE_PATH = os.Getenv("LOG_FILE_PATH")

	// Get certificate information from environment variables
	CERT_KEY_PATH   = os.Getenv("CERTIFICATE_KEY_PATH")
	CERT_CHAIN_PATH = os.Getenv("CERTIFICATE_CHAIN_PATH")

	ENDPOINT_HANDLER_MAP = map[string]func(http.ResponseWriter, *http.Request){
		"/":                         indexHandler,
		"blog." + DOMAIN_NAME + "/": blogHandler,
	}
)

func blogHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello from the blog handler!")
}

func indexHandler(writer http.ResponseWriter, request *http.Request) {
	http.ServeFile(writer, request, PAGES_DIR + "/index/index.html")
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

	log.Printf("Handling traffic for skippola.com on port%s", TLS_PORT)
	//err = http.ListenAndServeTLS(TLS_PORT, CERT_CHAIN_PATH, CERT_KEY_PATH, nil)
	err = http.ListenAndServe(":8077", nil)
	log.Fatal(err)
}
