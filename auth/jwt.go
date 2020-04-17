package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/sinistra/ecommerce-api/domain"
	"github.com/sinistra/ecommerce-api/utils"
)

func GenerateToken(user domain.User) (string, error) {
	var err error
	secret := os.Getenv("JWT_SECRET")
	issuer := os.Getenv("JWT_ISSUER")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Email,
		"iss":      issuer,
	})

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		log.Fatal(err)
	}

	return tokenString, nil
}

func TokenVerify(c *gin.Context) (int, string) {
	secret := os.Getenv("JWT_SECRET")
	authHeader := c.Request.Header.Get("Authorization")
	bearerToken := strings.Split(authHeader, " ")

	if len(bearerToken) == 2 {
		authToken := bearerToken[1]

		token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("there was an error")
			}

			return []byte(secret), nil
		})

		if err != nil {
			return http.StatusUnauthorized, err.Error()
		}

		if token.Valid {
			return http.StatusOK, ""
		} else {
			return http.StatusUnauthorized, "invalid token."
		}
	} else {
		return http.StatusUnauthorized, "token required"
	}
}

func JWTVerifyMiddleWare(c *gin.Context) {
	secret := os.Getenv("JWT_SECRET")
	authHeader := c.Request.Header.Get("Authorization")
	bearerToken := strings.Split(authHeader, " ")

	if len(bearerToken) == 2 {
		authToken := bearerToken[1]

		token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("there was an error")
			}

			return []byte(secret), nil
		})

		if err != nil {
			utils.RespondWithError(c, http.StatusUnauthorized, err.Error())
			return
		}

		if token.Valid {
			tokenUser := DecodeToken(c)
			c.Set("username", tokenUser)
			c.Next()
		} else {
			utils.RespondWithError(c, http.StatusUnauthorized, "invalid token.")
			return
		}
	} else {
		utils.RespondWithError(c, http.StatusUnauthorized, "token required")
		return
	}
}

func DecodeToken(c *gin.Context) string {

	secret := os.Getenv("JWT_SECRET")
	authHeader := c.Request.Header.Get("Authorization")
	bearerToken := strings.Split(authHeader, " ")

	authToken := ""
	if len(bearerToken) == 2 {
		authToken = bearerToken[1]
	} else {
		return "unknown"
	}

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(authToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	// ... error handling
	if err != nil {
		log.Println(err)
	}
	// log.Println(token)

	// do something with decoded claims
	for key, val := range claims {
		// fmt.Printf("Key: %v, value: %v\n", key, val)
		if key == "username" {
			return fmt.Sprintf("%v", val)
		}
	}
	return "unknown"
}

func Validate(username, password string) bool {
	return true
}
