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

// Login validates credentials and returns a JWT to be used on subsequent api calls
// @Summary Login a user
// @Description returns a JWT to be used on subsequent api calls
// @Accept  json
// @Produce  json
// @Param   username     body    string     true        "UserId"
// @Param   password     body    string     true        "Password"
// @Success 200 {object} domain.JWT	"ok"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Unauthorized"
// @Router /login [post]
func (u LoginController) Login(c *gin.Context) {
	var user domain.User
	var jwt domain.JWT

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "login failed binding.", "error": err.Error()})
	}

	if user.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "email is missing.", "data": user})
		return
	}
	if user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "password is missing.", "data": user})
		return
	}

	// ok := auth.LdapValidate(user.Username, user.Password)
	ok := auth.Validate(user.Email, user.Password)

	if ok == false {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "authentication failed", "data": user})
		return
	}

	token, err := auth.GenerateToken(user)

	if err != nil {
		log.Fatal(err)
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
