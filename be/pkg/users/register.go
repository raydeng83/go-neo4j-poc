package users

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type UserRegistration struct {
	User User `json:"user"`
}

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

type UserRegistrationHandler struct {
	Path           string
	UserRepository UserRepository
}

func (u *UserRegistrationHandler) Register(c *gin.Context) {
	jsonData, err := c.GetRawData()
	urRequest := UserRegistration{}
	err = json.Unmarshal(jsonData, &urRequest)
	if err != nil {
		panic(err)
	}
	rUser := urRequest.User
	err = u.UserRepository.RegisterUser(&rUser)
	if err != nil {
		panic(err)
	}

	c.JSON(201, gin.H{
		"user": User{
			Username: rUser.Username,
			Email:    rUser.Email,
		},
	})
}
