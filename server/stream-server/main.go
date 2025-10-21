package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jatin-jamdagni/fomo/server/stream-server/controllers"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello from fomo server")
	})

	router.GET("/movies", controllers.GetMovies())
	router.GET("/movie/:imdb_id", controllers.GetMovie())

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Failed to start server", err)
	}

}
