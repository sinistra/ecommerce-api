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

// ItemController is a struct that provides the controller vehicle
type ItemController struct{}

// GetItems returns a slice of items from db
func (s ItemController) GetItems(c *gin.Context) {

	request := make(map[string]string)
	// request["firstname"] = c.DefaultQuery("firstname", "")

	// log.Println(request)

	items, err := service.ItemsService.GetItems(request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": items})
}

func (s ItemController) GetItem(c *gin.Context) {
	// var item domain.Item
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err})
		return
	}

	item, err := service.ItemsService.GetItem(id)
	if err != nil {
		if err.Error() == "not found" {
			msg := fmt.Sprintf("%s not found.", id)
			c.JSON(http.StatusNotFound, gin.H{"message": msg})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": item})
}

func (s ItemController) AddItem(c *gin.Context) {
	var item domain.Item

	if err := c.BindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "service failed binding.", "error": err.Error()})
	}

	if item.Code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Enter missing fields"})
		return
	}

	itemID, err := service.ItemsService.AddItem(item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": itemID})
}

func (s ItemController) UpdateItem(c *gin.Context) {
	var item domain.Item

	if err := c.BindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "item failed binding.", "error": err.Error()})
	}

	if item.Id < 1 || item.Code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "missing fields"})
		return
	}

	count, err := service.ItemsService.UpdateItem(item)
	log.Println("items updated", count)

	if err != nil {
		if err.Error() == "not found" {
			msg := fmt.Sprintf("%d not found.", item.Id)
			c.JSON(http.StatusNotFound, gin.H{"message": msg, "error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("item %d updated", item.Id)})
}

func (s ItemController) RemoveItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "id is not a number", "error": err.Error()})
		return
	}

	count, err := service.ItemsService.RemoveItem(id)
	log.Println("items deleted", count)

	if err != nil {
		if err.Error() == "not found" {
			msg := fmt.Sprintf("%s not found.", id)
			c.JSON(http.StatusNotFound, gin.H{"message": msg, "error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("item %d removed.", id)})
}
