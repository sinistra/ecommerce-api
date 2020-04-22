package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/sinistra/ecommerce-api/auth"
	"github.com/sinistra/ecommerce-api/domain"
)

// LoginController is a struct that provides the controller vehicle
type LoginController struct{}

func (u LoginController) Login(c *gin.Context) {
	var loginUser domain.LoginRequest
	var jwt domain.JWT

	if err := c.BindJSON(&loginUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "login failed binding.", "error": err.Error()})
		return
	}

	if loginUser.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "username is missing.", "data": loginUser})
		return
	}
	if loginUser.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "password is missing.", "data": loginUser})
		return
	}

	ok, err := auth.Validate(loginUser.Username, loginUser.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "validation failed", "data": loginUser})
		return
	}

	if ok == false {
		log.Println("authentication failed", loginUser)
		loginUser.Password = ""
		c.JSON(http.StatusUnauthorized, gin.H{"message": "authentication failed", "data": loginUser})
		return
	}

	token, err := auth.GenerateToken(loginUser)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"message": "authentication failed", "data": loginUser})
	}

	jwt.Token = token
	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": jwt})
}

// CheckForToken checks if a JWT token exists in the payload
func (u LoginController) CheckForToken(c *gin.Context) {

	status, msg := auth.TokenVerify(c)

	if status == http.StatusOK {
		output := auth.DecodeToken(c)
		c.JSON(http.StatusOK, gin.H{"message": "ok", "data": output})
	} else {
		c.JSON(status, gin.H{"message": msg})
	}
}

// TestAuth checks if username is defined in the current Context
func (u LoginController) TestAuth(c *gin.Context) {
	username, ok := c.Get("username")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "username not in context"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": username})
}
