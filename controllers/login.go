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

func (u LoginController) Register(c *gin.Context) {
	var loginUser domain.LoginRequest

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

	user, err := service.UsersService.GetUserByEmail(loginUser.Username)
	if err != nil {
		log.Println(err)
		if err.Error() != "sql: no rows in result set" {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "cannot read user.", "data": loginUser})
			return
		}
	}

	if user.Id > 0 {
		msg := "User already exists."
		loginUser.Password = ""
		log.Println(msg)
		c.JSON(http.StatusBadRequest, gin.H{"message": msg, "data": loginUser})
		return
	}

	// spew.Dump(user)
	encryptedPassword := utils.EncryptPassword(loginUser.Password)
	uuidV4 := uuid.NewV4()
	uuidAsString := fmt.Sprintf("%s", uuidV4)
	// log.Println(uuidAsString)
	newUser := domain.User{
		Email:    loginUser.Username,
		Password: encryptedPassword,
		Status:   "unverified",
		UUID:     uuidAsString,
	}
	result, err := service.UsersService.AddUser(newUser)
	if err != nil {
		log.Println(err)
	}
	log.Println(result)
	newUser.Id = result

	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": result})
}

func (u LoginController) Verify(c *gin.Context) {
	uuidString := c.Param("id")

	user, err := service.UsersService.GetUserByUUID(uuidString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.UUID = ""
	user.Status = "Verified"
	result, err := service.UsersService.UpdateUser(user)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "server error", "error": err.Error()})
		return
	}
	c.String(http.StatusAccepted, "Record %d updated: %d", user.Id, result)
}
