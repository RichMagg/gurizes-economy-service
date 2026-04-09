package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.Writer.Write([]byte("Pong!"))
	})

	server.Run(":8080")
}
