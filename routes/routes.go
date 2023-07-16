package routes

import (
	"github.com/Mdromi/go-ecom-yt.git/controllers"
	"github.com/gin-gonic/gin"
)

// UserRoutes sets up the user-related routes
func UserRoutes(router *gin.Engine) {
	router.POST("/users/signup", controllers.Signup())
	router.POST("/users/login", controllers.Login())
	router.POST("/admin/addproduct", controllers.ProductViewerAdmin())
	router.GET("/users/productview", controllers.SearchProduct())
	router.GET("/users/search", controllers.SearchProductByQuery())
}
