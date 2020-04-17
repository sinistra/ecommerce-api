package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/sinistra/ecommerce-api/domain"
	"github.com/sinistra/ecommerce-api/service"
)

// UserController is a struct that provides the controller vehicle
type UserController struct{}

// Services is a slice of domain that controller functions will populate
// var Items []domain.Item

func (s UserController) GetUsers(c *gin.Context) {

	request := make(map[string][]string)
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "items failed binding.", "error": err.Error()})
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
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err})
		return
	}

	user, err := service.UsersService.GetUser(id)
	if err != nil {
		if err.Error() == "not found" {
			msg := fmt.Sprintf("%s not found.", id)
			c.JSON(http.StatusNotFound, gin.H{"message": msg})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": user})
}

func (s UserController) AddUser(c *gin.Context) {
	var user domain.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "user failed binding.", "error": err.Error()})
	}

	if user.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email is missing"})
		return
	}

	userID, err := service.UsersService.AddUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": userID})
}

func (s UserController) UpdateUser(c *gin.Context) {
	var user domain.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "user failed binding.", "error": err.Error()})
	}

	if user.Id < 1 || user.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "missing fields"})
		return
	}

	count, err := service.ItemsService.UpdateItem(user)
	log.Println("update count", count)

	if err != nil {
		if err.Error() == "not found" {
			msg := fmt.Sprintf("%d not found.", user.Id)
			c.JSON(http.StatusNotFound, gin.H{"message": msg, "error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%s updated", user.Id)})
}

func (s UserController) RemoveUser(c *gin.Context) {
	id := c.Param("id")

	count, err := service.UsersService.RemoveUser(id)

	if err != nil {
		if err.Error() == "not found" {
			msg := fmt.Sprintf("%s not found.", id)
			c.JSON(http.StatusNotFound, gin.H{"message": msg, "error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%s removed.", id)})
}
