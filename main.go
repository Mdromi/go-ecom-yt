package main

import (
	"log"
	"os"

	"github.com/Mdromi/go-ecom-yt.git/controllers"
	"github.com/Mdromi/go-ecom-yt.git/database"
	"github.com/Mdromi/go-ecom-yt.git/middleware"
	"github.com/Mdromi/go-ecom-yt.git/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

/**
* Main function for the application
*
* This function serves as the entry point of the application.
* It initializes the necessary components, sets up the server, and starts listening for incoming requests.
*
* Steps:
* 1. Load environment variables from the .env file.
* 2. Get the port number from the environment variables or use a default value.
* 3. Create an instance of the application with the necessary database collections.
* 4. Create a new Gin router and enable logging middleware.
* 5. Define the routes for user-related operations.
* 6. Apply authentication middleware to the router to protect certain routes.
* 7. Define the routes for cart management, address operations, and order placement.
* 8. Start the server and listen on the specified port.
*
* Diagrams:
* - Database Model Diagram
* - File Relationship Diagram
* - Code Update
* - API Documentation
* - Apps Documentation
 */
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	// Create an instance of the application with the necessary database collections
	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), (database.UserData(database.Client, "Users")))

	// Create a new Gin router and enable logging middleware
	router := gin.New()
	router.Use(gin.Logger())

	// Define the routes for user-related operations
	routes.UserRoutes(router)

	// Apply authentication middleware to the router to protect certain routes
	router.Use(middleware.Authentication())

	// Define the routes for cart management, address operations, and order placement
	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/listcart", controllers.GetItemFromCart())
	router.POST("/addaddress", controllers.AddAddress())
	router.PUT("/edithomeaddress", controllers.EditHomeAddress())
	router.PUT("/editworkaddress", controllers.EditWorkAddress())
	router.GET("/deleteaddresses", controllers.DeleteAddress())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	// Start the server and listen on the specified port
	log.Fatal(router.Run(":" + port))
}
