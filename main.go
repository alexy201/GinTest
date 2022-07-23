// Recipes API
//
// This is a sample recipes API. You can find out more about the API at https://github.com/PacktPublishing/Building-Distributed-Applications-in-Gin.
//
//	Schemes: http
//  Host: localhost:8080
//	BasePath: /
//	Version: 1.0.0
//	Contact: Mohamed Labouardy <mohamed@labouardy.com> https://labouardy.com
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
	"fmt"
	"log"
	"os"

	handlers "github.com/alexy201/GinTest/handlers"
	"github.com/gin-contrib/sessions"
	redisStore "github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var recipesHandler *handlers.RecipesHandler
var authHandler *handlers.AuthHandler

func init() {
	//openssl rand -base64 12 | docker secret create mongodb_password -
	//docker service create -d --name mongodb -e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD_FILE=/run/secrets/mongodb_password -p 27017:27017 mongo:latest
	//REDIS_URI="localhost:6379" JWT_SECRET=eUbP9shywUygMx7u MONGO_URI="mongodb://localhost:27017/test?authSource=admin" MONGO_DATABASE=demo go run main.go
	//docker run -d -v `pwd`/local-redis-stack.conf:/redis-stack.conf --name redis -p 6379:6379 redis:latest
	//ab -n 2000 -c 100 http://localhost:8080/recipes
	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URI"),
		Password: "",
		DB:       0,
	})
	status := redisClient.Ping()
	fmt.Println(status)
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err = client.Ping(context.TODO(),
		readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection("recipes")
	collectionUsers := client.Database(os.Getenv("MONGO_DATABASE")).Collection("users")
	authHandler = handlers.NewAuthHandler(ctx, collectionUsers)
	recipesHandler = handlers.NewRecipesHandler(ctx, collection, redisClient)
	log.Println("Connected to MongoDB")
}

func main() {
	router := gin.Default()

	store, _ := redisStore.NewStore(10, "tcp", os.Getenv("REDIS_URI"), "", []byte("secret"))
	router.Use(sessions.Sessions("recipes_api", store))

	authorized := router.Group("/")
	authorized.Use(authHandler.AuthMiddleware())
	{
		authorized.POST("/recipes", recipesHandler.NewRecipeHandler)
		authorized.PUT("/recipes/:id", recipesHandler.UpdateRecipeHandler)
		authorized.DELETE("/recipes/:id", recipesHandler.DeleteRecipeHandler)
		authorized.GET("/recipes/:id", recipesHandler.GetOneRecipeHandler)
	}
	router.GET("/recipes", recipesHandler.ListRecipesHandler)
	router.GET("/recipes/search", recipesHandler.SearchRecipesHandler)
	router.POST("/refresh", authHandler.RefreshHandler)
	router.POST("/signin", authHandler.SignInHandler)
	router.POST("/signup", authHandler.SignUpHandler)
	router.Run()
}
