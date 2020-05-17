package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/sinistra/ecommerce-api/domain"
	"github.com/sinistra/ecommerce-api/service"
	"github.com/sinistra/ecommerce-api/utils"
)

// UserController is a struct that provides the controller vehicle
type UserController struct{}

// Services is a slice of domain that controller functions will populate

func (s UserController) GetUsers(c *gin.Context) {

	request := make(map[string][]string)
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "items failed binding.", "error": err.Error()})
		return
	}
	log.Println(request)

	users, err := service.UsersService.GetUsers(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": users})
}

func (s UserController) GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}

	user, err := service.UsersService.GetUser(id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			msg := fmt.Sprintf("record %d not found.", id)
			c.JSON(http.StatusNotFound, gin.H{"message": msg})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Server error", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": user})
}

func (s UserController) AddUser(c *gin.Context) {
	var user domain.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "user failed binding.", "error": err.Error()})
		return
	}

	if user.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email is missing"})
		return
	}

	if len(user.Password) > 0 {
		user.Password = utils.EncryptPassword(user.Password)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Password is missing"})
		return
	}

	userID, err := service.UsersService.AddUser(user)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate") {
			c.JSON(http.StatusBadRequest, gin.H{"message": "A record with this email already exists"})
			return
		}
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": userID})
}

func (s UserController) UpdateUser(c *gin.Context) {
	var user domain.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "user failed binding.", "error": err.Error()})
		return
	}

	if user.Id < 1 || user.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "missing fields"})
		return
	}

	count, err := service.UsersService.UpdateUser(user)
	if err != nil {
		log.Println(err)
	}
	log.Println("update count", count)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%d updated", user.Id)})
}
func (s UserController) UpdatePassword(c *gin.Context) {
	// var user domain.User
	request := make(map[string]string)

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "post failed binding.", "error": err.Error()})
		return
	}
	log.Println(request)
	password, ok := request["password"]
	if !ok {
		msg := "password is not in the payload"
		log.Println(msg)
		c.JSON(http.StatusBadRequest, gin.H{"message": msg})
		return
	}

	username, ok := request["username"]
	if !ok {
		msg := "username is not in the payload"
		log.Println(msg)
		c.JSON(http.StatusBadRequest, gin.H{"message": msg})
		return
	}
	user, err := service.UsersService.GetUserByEmail(username)
	if err != nil {
		msg := fmt.Sprintf("could not find a record for %s", username)
		log.Println(err)
		c.JSON(http.StatusPreconditionFailed, gin.H{"message": msg})
		return
	}
	user.Password = utils.EncryptPassword(password)
	count, err := service.UsersService.UpdatePassword(user)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	log.Println("update count", count)
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("password updated")})
}

func (s UserController) RemoveUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "id is not a number", "error": err.Error()})
		return
	}
	count, err := service.UsersService.RemoveUser(id)
	log.Println("users deleted", count)

	if err != nil {
		if err.Error() == "not found" {
			msg := fmt.Sprintf("%s not found.", id)
			c.JSON(http.StatusNotFound, gin.H{"message": msg, "error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%d removed.", id)})
}
