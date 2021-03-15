package recipesHandler

import (
	"net/http"
)

// Handle all requests for specific recipe pages for any URL beginning with the
// path "/recipes/"
func recipesPageHandler(write http.ResponseWriter, request *http.Request) {

}

// Handle all requests for the recipe index page at "/recipes"
func recipesIndexHandler(writer http.ResponseWriter, request *http.Request) {

}
