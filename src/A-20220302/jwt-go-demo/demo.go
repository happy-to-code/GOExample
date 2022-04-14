package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

type LoginClaims struct {
	ID       int64
	Username string
	jwt.StandardClaims
}

func main() {
	token, err := GenerateToken()
	if err != nil {
		panic(err)
	}
	fmt.Println(token)
	fmt.Println("----------------------")
	// token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MTAwMDEsIlVzZXJuYW1lIjoiZnJhbmsiLCJleHAiOjE2NDg1MzQ2NTMsImlzcyI6ImZyYW5rIn0.OrsPsMWsQbVkRFHQcTV1fHn1kG3gDprMo75BmlwlXQw"
	claims, err := ParseToken(token)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", claims)

}

func GenerateToken() (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(60 * time.Second)
	issuer := "frank"

	claims := LoginClaims{
		ID:       10001,
		Username: "frank",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    issuer,
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("golang"))
	return token, err
}

func ParseToken(token string) (*LoginClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &LoginClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("golang"), nil
	})
	if err != nil {
		return nil, err
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*LoginClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// A sample use
var user = User{
	ID:       1,
	Username: "username",
	Password: "password",
}

func Login(c *gin.Context) {
	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	// compare the user from the request, with the one we defined:
	if user.Username != u.Username || user.Password != u.Password {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}
	token, err := CreateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}
	c.JSON(http.StatusOK, token)
}

func CreateToken(userid uint64) (string, error) {
	var err error
	// Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") // this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}
