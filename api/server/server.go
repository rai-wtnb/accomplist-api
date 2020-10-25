package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rai-wtnb/accomplist-api/controllers"
)

func Init() {
	r := router()
	r.Run()
}

func router() *gin.Engine {
	r := gin.Default()

	u := r.Group("/users")
	{
		ctrl := controllers.UserController{}
		u.GET("", ctrl.Index)
		u.POST("", ctrl.Create)
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
