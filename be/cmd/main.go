package main

import (
	"github.com/gin-gonic/gin"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/raydeng83/go-neo4j-poc/pkg/users"
	"log"
)

func main() {
	neo4jUri := "bolt://localhost:7687"
	neo4jUsername := "neo4j"
	neo4jPassword := "password"

	usersRepository := users.UserNeo4jRepository{
		Driver: driver(neo4jUri, neo4j.BasicAuth(neo4jUsername, neo4jPassword, "")),
	}
	registrationHandler := &users.UserRegistrationHandler{
		Path:           "/users",
		UserRepository: &usersRepository,
	}
	loginHandler := &users.UserLoginHandler{
		Path:           "/users/login",
		UserRepository: &usersRepository,
	}

	mux := gin.Default()
	mux.Any(registrationHandler.Path, registrationHandler.Register)
	mux.Any(loginHandler.Path, loginHandler.Login)

	err := mux.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}

func driver(target string, token neo4j.AuthToken) neo4j.Driver {
	result, err := neo4j.NewDriver(target, token)
	if err != nil {
		panic(err)
	}
	return result
}
