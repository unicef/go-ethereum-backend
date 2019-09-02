package userhandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qjouda/dignity-platform/backend/auth"
	"github.com/qjouda/dignity-platform/backend/datatype"
)

// SessionGet get session route
func SessionGet(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("user").(*datatype.User)
		c.JSON(http.StatusOK, user)
	}
}

// SessionDestroy logout route
func SessionDestroy(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth.Logout(c)
		c.JSON(http.StatusOK, nil)
	}
}
