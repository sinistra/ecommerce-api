package controllers

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/sinistra/ecommerce-api/domain"
	"github.com/sinistra/ecommerce-api/service"
	"github.com/sinistra/ecommerce-api/utils"
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
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if err := c.Bind(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "service failed binding.", "error": err.Error()})
		return
	}

	if item.Code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Enter missing fields"})
		return
	}

	imagePtr, err := processImage(c, item)
	if err != nil {
		msg := "cannot save image"
		log.Println(msg)
		c.JSON(http.StatusInternalServerError, gin.H{"message": msg, "error": err.Error()})
		return
	}

	item.Image = *imagePtr

	itemID, err := service.ItemsService.AddItem(item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	item, err = service.ItemsService.GetItem(itemID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": item})
}

func (s ItemController) UpdateItem(c *gin.Context) {
	var item domain.Item
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if err := c.Bind(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "item failed binding.", "error": err.Error()})
		return
	}

	if item.Id < 1 || item.Code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "missing fields"})
		return
	}

	imagePtr, err := processImage(c, item)
	if err != nil {
		msg := "cannot save image"
		log.Println(msg)
		c.JSON(http.StatusInternalServerError, gin.H{"message": msg, "error": err.Error()})
		return
	}

	item.Image = *imagePtr

	// spew.Dump(item)
	count, err := service.ItemsService.UpdateItem(item)
	if err != nil {
		if err.Error() == "not found" {
			msg := fmt.Sprintf("%d not found.", item.Id)
			c.JSON(http.StatusNotFound, gin.H{"message": msg, "error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Server error"})
		return
	}
	log.Println("items updated", count)

	item, err = service.ItemsService.GetItem(item.Id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("item %d updated", item.Id), "data": item})
}

func processImage(c *gin.Context, item domain.Item) (*string, error) {
	file, header, err := c.Request.FormFile("upload")
	var imagePath *string
	if err != nil {
		log.Println("image upload empty or there was an error.", err)
		return nil, err
	} else {
		imagePath, err = saveImage(file, header, item.Code)
		if err != nil {
			msg := "cannot save image"
			log.Println(msg, err)
			// c.JSON(http.StatusInternalServerError, gin.H{"message": msg, "error": err.Error()})
			return nil, err
		}
		// item.Image = *imagePath
	}
	return imagePath, nil
}

func saveImage(file multipart.File, header *multipart.FileHeader, itemCode string) (*string, error) {
	filename := header.Filename
	// log.Println(filename)
	fileParts := strings.Split(filename, ".")
	// spew.Dump(fileParts)

	utils.CreateDirIfNotExist("public/images")
	log.Println("Upload successful")

	imagePath := itemCode + "." + fileParts[1]
	out, err := os.Create("./public/images/" + imagePath)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &imagePath, nil
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
