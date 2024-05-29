package server

import (
	"mmr/backend/controllers"
	"mmr/backend/docs"
	"mmr/backend/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := router.Group("/api/v1")
	{
		mg := v1.Group("/mmr", middleware.RequireAuth)
		{
			match := new(controllers.MatchController)
			mg.POST("/matches", match.SubmitMatch)
			mg.GET("/matches", match.GetMatches)
		}
		s := v1.Group("/stats", middleware.RequireAuth)
		{
			leaderboard := new(controllers.LeaderboardController)
			s.GET("/leaderboard", leaderboard.GetLeaderboard)
		}
		a := v1.Group("/admin", middleware.RequireAuth)
		{
			admin := new(controllers.AdminController)
			a.POST("/recalculate", admin.RecalculateMatches)
		}
	}

	router.GET("/swagger", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusPermanentRedirect, "/swagger/index.html")
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router

}
