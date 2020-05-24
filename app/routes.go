package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"

	"github.com/sinistra/ecommerce-api/auth"
	"github.com/sinistra/ecommerce-api/controllers"
)

var (
	userController  = controllers.UserController{}
	loginController = controllers.LoginController{}
	itemController  = controllers.ItemController{}
)

func mapUrls() {
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong "+fmt.Sprint(time.Now().Unix()))
	})

	router.POST("/login", loginController.Login)
	router.POST("/auth/register", loginController.Register)
	router.POST("/auth/update", userController.UpdatePassword)
	router.GET("/auth/verify/:id", loginController.Verify)
	router.GET("/checktoken", loginController.CheckForToken)
	router.GET("/testauth", auth.JWTVerifyMiddleWare, loginController.TestAuth)

	api := router.Group("/api", auth.JWTVerifyMiddleWare)

	api.GET("/users", userController.GetUsers)
	api.GET("/users/:id", userController.GetUser)
	api.POST("/users", userController.AddUser)
	api.PUT("/users", userController.UpdateUser)
	api.DELETE("/users/:id", userController.RemoveUser)

	// the 1st 2 routes are unsecured
	router.POST("/items", itemController.GetItems)
	router.GET("/items/:id", itemController.GetItem)
	api.POST("/items", itemController.AddItem)
	api.PUT("/items", itemController.UpdateItem)
	api.DELETE("/items/:id", itemController.RemoveItem)
}
