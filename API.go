package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

var users = []User{
    {ID: 1, Name: "Budi"},
    {ID: 2, Name: "Ani"},
}

func main() {
    r := gin.Default()
    
    // GET all users
    r.GET("/users", func(c *gin.Context) {
        c.JSON(http.StatusOK, users)
    })
    
    // GET user by id
    r.GET("/users/:id", func(c *gin.Context) {
        id := c.Param("id")
        c.JSON(http.StatusOK, gin.H{
            "message": "User ID " + id,
        })
    })
    
    // POST create user
    r.POST("/users", func(c *gin.Context) {
        var newUser User
        if err := c.BindJSON(&newUser); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        users = append(users, newUser)
        c.JSON(http.StatusCreated, newUser)
    })
    
    r.Run(":8090")
}
