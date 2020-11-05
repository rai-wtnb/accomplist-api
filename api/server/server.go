package server

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"github.com/rai-wtnb/accomplist-api/controllers"
)

// Init starts router.
func Init() {
	r := router()
	r.Run()
}

func router() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/ids", controllers.IndexID)

	u := r.Group("/users")
	{
		ctrl := controllers.UserController{}
		u.GET("", ctrl.Index)
		u.POST("/signup", ctrl.Signup)
		u.POST("/login", ctrl.Login)
		u.POST("/logout", ctrl.Logout)
		u.GET("/:id", ctrl.Show)
		u.PUT("/:id", ctrl.Update)
		u.DELETE("/:id", ctrl.Delete)
	}

	// p := r.Group("/lists")
	// {
	//     ctrl := controllers.ListController{}
	//     p.GET("", ctrl.Index)
	//     p.POST("", ctrl.Create)
	//     p.GET("/:id", ctrl.Show)
	//     p.PUT("/:id", ctrl.Update)
	//     p.DELETE("/:id", ctrl.Delete)
	// }

	return r
}
