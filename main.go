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

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), (database.UserData(database.Client, "Users")))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantBuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}
