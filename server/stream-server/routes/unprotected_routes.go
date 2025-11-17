package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jatin-jamdagni/fomo/server/stream-server/controllers"
)

func SetupUnProtectedRoutes(router *gin.Engine) {

	router.GET("/movies", controllers.GetMovies())
	router.POST("/register", controllers.RegisterUser())
	router.POST("/login", controllers.LoginUser())

}
