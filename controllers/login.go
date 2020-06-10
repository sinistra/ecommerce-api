package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"

	"github.com/sinistra/ecommerce-api/auth"
	"github.com/sinistra/ecommerce-api/domain"
	"github.com/sinistra/ecommerce-api/service"
	"github.com/sinistra/ecommerce-api/utils"
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

	user, err := auth.Validate(loginUser.Username, loginUser.Password)
	if err != nil {
		loginUser.Password = ""
		if err.Error() == "user failed validation" {
			log.Println("authentication failed", loginUser)
			c.JSON(http.StatusUnauthorized, gin.H{"message": "authentication failed", "data": loginUser})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"message": "validation failed", "data": loginUser})
		return
	}

	token, err := auth.GenerateToken(loginUser)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"message": "authentication failed", "data": loginUser})
	}

	jwt.Token = token
	jwt.UserName = user.Email
	jwt.FirstName = user.FirstName
	jwt.LastName = user.LastName
	jwt.Status = user.Status
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

func (u LoginController) Register(c *gin.Context) {
	var newUser domain.User

	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "login failed binding.", "error": err.Error()})
		return
	}

	if newUser.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "email is missing.", "data": newUser})
		return
	}
	if newUser.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "password is missing.", "data": newUser})
		return
	}

	existingUser, err := service.UsersService.GetUserByEmail(newUser.Email)
	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			log.Println(err)
			newUser.Password = ""
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error finding user.", "data": newUser})
			return
		}
	}

	if existingUser.Id > 0 {
		msg := "User already exists."
		newUser.Password = ""
		log.Println(msg)
		c.JSON(http.StatusBadRequest, gin.H{"message": msg, "data": newUser})
		return
	}

	// spew.Dump(newUser)
	encryptedPassword := utils.EncryptPassword(newUser.Password)
	uuidV4 := uuid.NewV4()
	uuidAsString := fmt.Sprintf("%s", uuidV4)
	// log.Println(uuidAsString)
	user := domain.User{
		Email:     newUser.Email,
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Password:  encryptedPassword,
		Status:    "unverified",
		UUID:      &uuidAsString,
	}

	// spew.Dump(user)
	result, err := service.UsersService.AddUser(user)
	if err != nil {
		log.Println(err)
		user.Password = ""
		c.JSON(http.StatusBadRequest, gin.H{"message": "error adding user to db", "data": user})
		return

	}
	// log.Println(result)
	user.Id = result
	user.Password = ""

	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": user})
}

func (u LoginController) Verify(c *gin.Context) {
	uuidString := c.Param("id")

	user, err := service.UsersService.GetUserByUUID(&uuidString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.UUID = nil
	user.Status = "verified"
	result, err := service.UsersService.UpdateUser(user)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "server error", "error": err.Error()})
		return
	}
	c.String(http.StatusAccepted, "Record %d updated: %d", user.Id, result)
}
