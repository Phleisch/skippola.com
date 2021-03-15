package recipesHandler

import (
	"json"
	"ioutil"
	"strings"
	"net/http"
)

// struct for a recipe; all fields are required and both ingredients and
// instructions lists must have at least one item each
type recipe struct {
	Name			string			`json:"name"`
	PrepTime		string			`json:"prepTime"`
	TotalTime		string			`json:"totalTime"`
	Ingredients		[]ingredient	`json:"ingredients"`
	Instructions	[]string		`json:"instructions"`
}

// struct for an ingredient in a recipe; name is required, amount and prep are
// optional
type ingredient struct {
	Name	string	`json:"name"`
	Amount	string	`json:"amount"`
	Prep	string	`json:"prep"`
}

func generateRecipePage(r recipe, pageTemplate, pathName string) {

}

// Handle all requests for specific recipe pages for any URL beginning with the
// path "/recipes/"
func PageHandler(write http.ResponseWriter, request *http.Request) {
	// get the recipe page to fetch from the URL
	recipeName := strings.TrimPrefix(request.URL.Path, "/recipes/")

	// open the file representing the recipe as json
	recipeJson, err := ioutil.ReadFile(recipeName + ".json")
	if err != nil { fmt.Println(err) }

	// unpack the json file into golang struct representation of the recipe
	var r recipe
	err = json.Unmarshal(recipeJson, &r)
	if err != nil { fmt.Println(err) }

	// open the recipe page template for formatting
	recipeWebpageTemplate, err := ioutil.ReadFile("recipeFormat.html")
	if err != nil { fmt.Println(err) }

	// format the recipe page template with specifics of the recipe to generate
	// the recipe's webpage
	generateRecipePage(r, string(recipeWebpageTemplate), recipeName)
}

// Handle all requests for the recipe index page at "/recipes"
func IndexHandler(writer http.ResponseWriter, request *http.Request) {

}
