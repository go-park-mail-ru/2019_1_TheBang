package main

import (
	"2019_1_TheBang/pkg/main-serivce-pkg/leaderboard"
	"2019_1_TheBang/pkg/main-serivce-pkg/login"
	"2019_1_TheBang/pkg/main-serivce-pkg/logout"
	"2019_1_TheBang/pkg/main-serivce-pkg/user"
	"2019_1_TheBang/pkg/public/middleware"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func setUpMainRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CorsMiddlewareGin,
		middleware.AuthMiddlewareGin,
		middleware.MetricMiddleware)

	router.POST("/auth", login.LogInHandler)
	router.DELETE("/auth", logout.LogoutHandler)

	router.OPTIONS("/auth", func(c *gin.Context) {})

	router.POST("/user", user.MyProfileCreateHandler)
	router.GET("/user", user.MyProfileInfoHandler)
	router.PUT("/user", user.MyProfileInfoUpdateHandler)
	router.POST("/user/avatar", user.ChangeProfileAvatarHandler)

	router.OPTIONS("/user", func(c *gin.Context) {})
	router.OPTIONS("/user/avatar", func(c *gin.Context) {})

	router.GET("/icon/:filename", user.GetIconHandler)

	router.GET("/leaderbord/:page", leaderboard.LeaderbordHandler)

	router.GET("/metrics", func(ctx *gin.Context) {
		h := promhttp.Handler()
		h.ServeHTTP(ctx.Writer, ctx.Request)
	})

	return router
}
