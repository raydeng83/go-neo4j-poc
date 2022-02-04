package users

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

type UserLogin struct {
	User User `json:"user"`
}

type LoggedInUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

type UserLoginHandler struct {
	Path           string
	UserRepository UserRepository
}

func (u *UserLoginHandler) Login(c *gin.Context) {
	jsonData, err := c.GetRawData()
	userLoginRequest := UserLogin{}
	_ = json.Unmarshal(jsonData, &userLoginRequest)
	rUser := userLoginRequest.User
	user, _ := u.UserRepository.FindByEmailAndPassword(
		rUser.Email,
		rUser.Password,
	)
	if user == nil {
		c.JSON(401, gin.H{})
		return
	}
	token, err := CreateToken(user)
	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"loggedInUser": LoggedInUser{
			Username: user.Username,
			Email:    user.Email,
			Token:    token,
		},
	})
}

func CreateToken(user *User) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user.Username
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("my-secret-key"))
}
