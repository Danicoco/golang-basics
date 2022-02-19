package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON((200), gin.H{
			"success": true,
			"data":    "Route successfully hit",
		})
	})

	router.GET("/params/:name", func(c *gin.Context) {
		name, err := c.GetRawData()

		if err != nil {
			fmt.Println("This is error")
		}

		fmt.Println(name)

		c.JSON((200), gin.H{
			"success": true,
			"data":    name,
		})
	})

	router.POST("/v1/hit", func(c *gin.Context) {
		name := c.PostForm("name")
		c.JSON((200), gin.H{
			"success": true,
			"data":    name,
		})
	})

	router.Run(":3000")
}
