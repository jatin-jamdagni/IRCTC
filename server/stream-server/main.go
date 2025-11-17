package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jatin-jamdagni/fomo/server/stream-server/routes"
)

func main() {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.String(200, "Hi, I'm healty")
	})

	routes.SetupUnProtectedRoutes(router)

	routes.SetupProtectedRoutes(router)

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Failed to start server", err)
	}

}
