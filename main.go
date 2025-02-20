package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ShardenduMishra22/GoLangToDoList/database"
	"github.com/ShardenduMishra22/GoLangToDoList/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection // Declare a global variable to hold the collection

func main() {
	fmt.Println("ToDo List Project!!")

	if os.Getenv("ENV") != "production" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file:", err)
		}
	}

	// Connect to the database and get the collection
	collection = database.ConnectToDatabase()
	
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://golangtodo-front.onrender.com", // Change to your frontend origins
		AllowMethods: "GET,POST,PATCH,DELETE,PUT",                           // Allowed methods
		AllowHeaders: "Origin, Content-Type, Accept",                  // Allowed headers
	}))
	
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message1": "Welcome to ToDo List Project",
			"message2": "This is a Sample Response to test if the application",
		})
	})

	routes.SetupRoutes(app, collection)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Fatal(app.Listen("0.0.0.0:" + port))
}
