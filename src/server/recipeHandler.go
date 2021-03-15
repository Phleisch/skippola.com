package recipesHandler

import (
	"strings"
	"net/http"
)

// Handle all requests for specific recipe pages for any URL beginning with the
// path "/recipes/"
func recipesPageHandler(write http.ResponseWriter, request *http.Request) {
	// get the recipe page to fetch from the URL
	recipeName := strings.TrimPrefix(request.URL.Path, "/recipes/")

	// open the file representing the recipe as json
	recipeJson, err := ioutil.ReadFile(recipeName + ".json")
	if err != nil { fmt.Println(err) }

	// unpack the json file into golang struct representation of the recipe
	var recipe Recipe
	err = json.Unmarshal(recipeJson, &recipe)
	if err != nil { fmt.Println(err) }

	// open the recipe page template for formatting
	recipeWebpageTemplate, err := ioutil.ReadFile("recipeFormat.html")
	if err != nil { fmt.Println(err) }

	// format the recipe page template with specifics of the recipe to generate
	// the recipe's webpage
	generateRecipePage(recipe, string(recipeWebpageTemplate), recipeName)
}

// Handle all requests for the recipe index page at "/recipes"
func recipesIndexHandler(writer http.ResponseWriter, request *http.Request) {

}
