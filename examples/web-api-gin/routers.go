package main

import "github.com/gin-gonic/gin"

func getRouter() {

}

func routers() {
	router := gin.Default()

	router.GET("/v1/user", getRouter())
}
