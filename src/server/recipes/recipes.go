package recipes

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const (
	RECIPES_DIR            = "../pages/recipes/"
	RECIPE_PAGE_TEMPLATE   = RECIPES_DIR + "recipePageTemplate.html"
	RECIPES_INDEX_TEMPLATE = RECIPES_DIR + "recipesIndexTemplate.html"
	RECIPE_FILE_EXTENSION  = ".json"
)

// struct for a recipe; all fields are required and both ingredients and
// instructions lists must have at least one item each
type recipe struct {
	Name         string `json:"name"`
	UrlName      string
	PrepTime     string       `json:"prepTime"`
	TotalTime    string       `json:"totalTime"`
	Ingredients  []ingredient `json:"ingredients"`
	Instructions []string     `json:"instructions"`
}

// struct for an ingredient in a recipe; name is required, amount and prep are
// optional
type ingredient struct {
	Name   string `json:"name"`
	Amount string `json:"amount"`
	Prep   string `json:"prep"`
}

// Handle all requests for specific recipe pages for any URL beginning with the
// path "/recipes/"
func PageHandler(writer http.ResponseWriter, request *http.Request) {
	// get the recipe page to fetch from the URL
	recipeName := strings.TrimPrefix(request.URL.Path, "/recipes/")

	// open the file representing the recipe as json
	recipeJson, err := ioutil.ReadFile(
		RECIPES_DIR + recipeName + RECIPE_FILE_EXTENSION)

	if err != nil {
		fmt.Println(err)
	}

	// unpack the json file into golang struct representation of the recipe
	var r recipe
	err = json.Unmarshal(recipeJson, &r)
	if err != nil {
		fmt.Println(err)
	}
	r.UrlName = recipeName

	// open the recipe page template for formatting
	pageTemplate, err := ioutil.ReadFile(RECIPE_PAGE_TEMPLATE)
	if err != nil {
		fmt.Println(err)
	}

	// open and parse the recipe template
	tmplt, err := template.New(recipeName).Parse(string(pageTemplate))
	if err != nil {
		fmt.Println(err)
	}

	// execute the template with specific recipe data to create the recipe page
	// and write as a response
	tmplt.Execute(writer, r)
}

// return all json recipe files from the directory RECIPES_DIR
func getRecipeFiles() []string {
	recipesDirEntries, err := os.ReadDir(RECIPES_DIR)
	if err != nil {
		fmt.Println(err)
	}

	var recipeFiles []string
	for _, recipesDirEntry := range recipesDirEntries {
		if recipesDirEntry.IsDir() {
			continue
		}
		if strings.HasSuffix(recipesDirEntry.Name(), RECIPE_FILE_EXTENSION) {
			recipeFiles = append(recipeFiles, recipesDirEntry.Name())
		}
	}

	return recipeFiles
}

// create and return a map of recipe url names to actual recipe names
func mapRecipeUrlToRecipeName() map[string]string {
	recipeFiles := getRecipeFiles()
	recipeUrlToName := make(map[string]string)

	for _, recipeFile := range recipeFiles {
		// open the file representing the recipe as json
		recipeJson, err := ioutil.ReadFile(RECIPES_DIR + recipeFile)
		if err != nil {
			fmt.Println(err)
		}

		// unpack the json file into golang struct representation of the recipe
		var r recipe
		err = json.Unmarshal(recipeJson, &r)
		if err != nil {
			fmt.Println(err)
		}

		recipeUrlToName[strings.TrimSuffix(recipeFile, RECIPE_FILE_EXTENSION)] = r.Name
	}

	return recipeUrlToName
}

// Handle all requests for the recipe index page at "/recipes"
func IndexHandler(writer http.ResponseWriter, request *http.Request) {
	recipeUrlToName := mapRecipeUrlToRecipeName()
	// open the recipe page template for formatting
	pageTemplate, err := ioutil.ReadFile(RECIPES_INDEX_TEMPLATE)
	if err != nil {
		fmt.Println(err)
	}

	// open and parse the recipe template
	tmplt, err := template.New("recipesIndex").Parse(string(pageTemplate))
	if err != nil {
		fmt.Println(err)
	}

	// execute the template with specific recipe data to create the recipe page
	// and write as a response
	tmplt.Execute(writer, recipeUrlToName)
}
