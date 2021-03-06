package server

import (
	"github.com/gin-contrib/cors"
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

	// CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"*",
	}
	r.Use(cors.New(config))

	u := r.Group("/users")
	{
		ctrl := controllers.UserController{}
		u.GET("", ctrl.Index)
		u.POST("/signup", ctrl.Signup)
		u.POST("/login", ctrl.Login)
		u.POST("/logout", ctrl.Logout)
		u.GET("/:id", ctrl.Show)
		u.PUT("/:id", ctrl.Update)
		u.PUT("/:id/img", ctrl.UpdateImg)
		u.DELETE("/:id", ctrl.Delete)
	}

	l := r.Group("/lists")
	{
		ctrl := controllers.ListController{}
		l.GET("", ctrl.Index)
		l.POST("", ctrl.Create)
		l.GET("/users/:id", ctrl.IndexByUserID)
		l.GET("/specific/:id", ctrl.Show)
		l.PUT("/specific/:id", ctrl.Update)
		l.DELETE("/specific/:id", ctrl.Delete)
	}

	f := r.Group("/feedbacks")
	{
		ctrl := controllers.FeedbackController{}
		f.GET("", ctrl.Index)
		f.POST("", ctrl.Create)
		f.GET("/:id", ctrl.Show)
		f.PUT("/:id", ctrl.Update)
		f.PUT("/:id/img", ctrl.UpdateImgByListID)
		f.DELETE("/:list-id", ctrl.Delete)
	}

	rel := r.Group("/relations")
	{
		ctrl := controllers.RelationController{}
		rel.GET("/follows/:id", ctrl.FollowIndex)
		rel.GET("/followers/:id", ctrl.FollowerIndex)
		rel.GET("/count/:id", ctrl.Count)
		rel.POST("/isfollow", ctrl.IsFollow)
		rel.POST("", ctrl.Create)
		rel.DELETE("", ctrl.Delete)
	}

	research := r.Group("/research")
	{
		ctrl := controllers.ResearchController{}
		research.GET("", ctrl.ResearchIndex)
	}

	return r
}
