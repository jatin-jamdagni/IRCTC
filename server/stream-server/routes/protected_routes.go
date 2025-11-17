package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jatin-jamdagni/fomo/server/stream-server/controllers"
	"github.com/jatin-jamdagni/fomo/server/stream-server/middleware"
)

func SetupProtectedRoutes(router *gin.Engine) {
	router.Use(middleware.AuthMiddleware())

	router.GET("/movie/:imdb_id", controllers.GetMovie())
	router.POST("/addmovie", controllers.AddMovie())

}
