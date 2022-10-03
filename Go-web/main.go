package main

import "github.com/gin-gonic/gin"

func saludoHandler(ctx *gin.Context) {
	ctx.JSON(299, gin.H{
		"message": "Hola\n",
	})
}

func main() {
	router := gin.Default()
	router.GET("/saludo", saludoHandler)
	router.Run()
}
