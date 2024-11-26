package routers

import (
	"web-gin/controllers"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("views/*")
	store := cookie.NewStore([]byte("loginuser"))
	router.Use(sessions.Sessions("mysession", store))
	{
		router.GET("/", controllers.HomeGet)
		router.GET("/register", controllers.RegisterGet)
		router.POST("/register", controllers.RegisterPost)

		router.GET("/login", controllers.LoginGet)
		router.POST("/login", controllers.LoginPost)

		router.GET("/exit", controllers.ExitGet)
	}

	return router
}
