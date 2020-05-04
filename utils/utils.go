package utils

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math"
	"net/http"
	"os"
	"strings"
)

func RespondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

func RespondWithJSON(w http.ResponseWriter, data interface{}) {
	_ = json.NewEncoder(w).Encode(data)
}

func DumpStructAsJson(object interface{}) {
	b, err := json.Marshal(object)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Paginate(page, pageSize, recordCount int) (low, high, pageNum, maxPage int) {
	page = Max(page, 1)
	pageSize = Max(pageSize, 5)
	pageSize = Min(pageSize, recordCount)

	maxPage = int(math.Ceil(float64(recordCount) / float64(pageSize)))
	pageNum = Min(page, maxPage)

	low = (pageNum - 1) * pageSize
	high = Min(pageNum*pageSize, recordCount)

	// spew.Dump(pageNum, pageSize, low, high, recordCount)
	return low, high, pageNum, maxPage
}

func Contains(s []string, t string) bool {
	for _, a := range s {
		if strings.ToLower(a) == strings.ToLower(t) {
			return true
		}
	}
	return false
}

func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}

func EncryptPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)
}

func MakeKey(keylen int) (string, error) {
	key := make([]byte, keylen)
	_, err := rand.Read(key)
	if err != nil {
		// handle error here
		log.Println(err)
		return "", err
	}

	log.Println(string(key))
	return fmt.Sprintf("%x", key), nil
}
