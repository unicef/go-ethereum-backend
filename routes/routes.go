package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/qjouda/dignity-platform/backend/assethandlers"
	"github.com/qjouda/dignity-platform/backend/commonhandlers"
	"github.com/qjouda/dignity-platform/backend/datatype"
	"github.com/qjouda/dignity-platform/backend/githandlers"
	"github.com/qjouda/dignity-platform/backend/routermiddleware"
	"github.com/qjouda/dignity-platform/backend/timelinehandlers"
	"github.com/qjouda/dignity-platform/backend/userhandlers"
	api "gopkg.in/appleboy/gin-status-api.v1"
)

//SetupRouting sets up routes
func SetupRouting(sc datatype.ServiceContainer) *gin.Engine {
	r := gin.Default()
	// staic
	r.Static("/static", "./public")
	r.Static("/img", "./public/img")
	// html
	web := r.Group("/")
	web.Use(routermiddleware.Session())
	web.Use(routermiddleware.SessionSetUser(sc.UserService))
	{
		commonhandlers.InjectHandlers(sc, web)
	}
	// Web apis
	wapi := r.Group("/wapi")
	wapi.Use(routermiddleware.Session())
	wapi.Use(routermiddleware.SessionSetUser(sc.UserService))
	wapi.Use(routermiddleware.CheckCsrfToken())
	{
		wapi.GET("/csrf", routermiddleware.SetCsrfToken())
		userhandlers.InjectHandlers(sc, wapi)
		assethandlers.InjectHandlers(sc, wapi)
		timelinehandlers.InjectHandlers(sc, wapi)
		githandlers.InjectHandlers(sc, wapi)
	}
	// API
	v1 := r.Group("/api")
	v1.Use(routermiddleware.HeadersNoCache())
	v1.Use(routermiddleware.HeadersCors())
	{
		v1.GET("/status", routermiddleware.APIAuth(), api.StatusHandler)
	}
	return r
}
