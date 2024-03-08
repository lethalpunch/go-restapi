package main

import (
	"errors"
	"fmt"
	"livehost"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	Name  string `json:"naam"`
	Age   int    `json:"umra"`
	Alive bool   `json:"jinda"`
}

var users = []user{
	{Name: "Amar", Age: 30, Alive: true},
	{Name: "Akbar", Age: 28, Alive: true},
	{Name: "Anthony", Age: 26, Alive: false},
}

func main() {

	for i := 1; i < 255; i++ {
		var livehosts []net.IP
		livehost, _ := livehost.SideroLiveHost(i)
		livehosts = append(livehosts, livehost)
		if len(livehosts) > 0 {
			fmt.Println("sidero hosts in network", livehosts)
		} else {
			fmt.Println("no sidero host in network")
		}
		//go sideroliveSideroLiveHost(i)
		// here need to get the list of all ips in network
		//time.Sleep(300000)
	}

	router := gin.Default()
	router.GET("/users", getUsers)
	router.GET("/users/:name", getUser)
	router.POST("/users", addUsers)
	router.Run("localhost:9090")
}

func getUsers(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, users)
}

func addUsers(context *gin.Context) {
	var newUser user
	if err := context.BindJSON(&newUser); err != nil {
		return
	}
	users = append(users, newUser)

	context.IndentedJSON(http.StatusCreated, newUser)
}

func getUserById(name string) (*user, error) {
	for i, u := range users {
		if u.Name == name {
			return &users[i], nil
		}
	}

	return nil, errors.New("User not found")
}

func getUser(context *gin.Context) {
	name := context.Param("name")
	user, err := getUserById(name)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not Found"})
	}
	context.IndentedJSON(http.StatusOK, user)
}
