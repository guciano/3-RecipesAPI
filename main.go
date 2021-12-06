// Recipes API
//
// This is a sample recipes API. You can find out more about the API at https://github.com/PacktPublishing/Building-Distributed-Applications-in-Gin.
//
//	Schemes: http
//  Host: localhost:8080
//	BasePath: /
//	Version: 1.0.0
//	Contact:  Adam Othasha Guciano <adamothasha@gmail.com> https://about.me/adamsgucianos17
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
// swagger:meta
package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	handlers "3-RecipesAPI/handlers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

/*Each recipe should have a name, a list of ingredients, a list of instructions or steps,
and a publication date. Moreover, each recipe belongs to a set of categories or tags (for example
vegan, Italian, pastry, salads, and so on), as well an ID, which is unique identifier to
differentiate each recipe in the database. Also specify the tags on each field using backtick
annotation; for example, `json:"NAME"`. This allows us to map each field to a different name when
we send them as response, since JSON and GO have different naming conventions.
*/

var recipesHandler *handlers.RecipesHandler

func init() {
	ctx := context.Background()
	client, err := mongo.Connect(ctx,
		options.Client().ApplyURI("mongodb://adminMongo:sembarang@localhost:27017/?authSource=admin&readPreference=primary&appname=MongoDB%20Compass&ssl=false"))
	if err = client.Ping(context.TODO(),
			readpref.Primary()) ; err != nil {
				log.Fatal(err)
			}
			log.Println("Terkoneksi ke MongoDB")
			
			collection := client.Database("DEMO_DATABASE").Collection("recipes")
			recipesHandler = handlers.NewRecipesHandler(ctx, collection)
			
}

func main() {
	router := gin.Default()
	router.POST("/recipes", recipesHandler.NewRecipeHandler)
	router.GET("/recipes", recipesHandler.ListRecipesHandler)
	router.PUT("/recipes/:id", recipesHandler.UpdateRecipeHandler)
	router.DELETE("/recipes/:id", recipesHandler.DeleteRecipeHandler)
	//router.GET("/recipes/search", SearchRecipesHandler)
	router.GET("/recipes/:id", recipesHandler.GetOneRecipeHandler)
	router.Run()
}