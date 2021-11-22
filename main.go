package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name string
	Age  int
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	router := gin.Default()

	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", handler)

	// router.Run(":3000")
	router.Run(":" + port)
}

func handler(ctx *gin.Context) {

	user := User{"User", 20}

	ctx.HTML(200, "index.html", gin.H{
		"user": user,
	})
}
