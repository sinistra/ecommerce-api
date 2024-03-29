package auth

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sinistra/ecommerce-api/domain"
	"github.com/sinistra/ecommerce-api/service"
	"github.com/sinistra/ecommerce-api/utils"
	"golang.org/x/crypto/bcrypt"
)

func GenerateToken(user domain.LoginRequest) (string, error) {
	var err error
	secret := os.Getenv("JWT_SECRET")
	issuer := os.Getenv("JWT_ISSUER")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
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

func Validate(username, password string) (*domain.User, error) {

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	passwordString := string(hashedPassword)
	user, err := service.UsersService.GetUserByEmail(username)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// TODO remove these comments
	log.Println(username, password)
	log.Println(passwordString, user.Password)

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) == nil {
		return &user, nil
	}

	return nil, fmt.Errorf("user failed validation")
}
