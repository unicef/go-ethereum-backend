package userhandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qjouda/dignity-platform/backend/auth"
	"github.com/qjouda/dignity-platform/backend/datatype"
)

//Logout logout route
func Logout(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth.Logout(c)
		//c.JSON(http.StatusOK, gin.H{})
		c.Redirect(http.StatusPermanentRedirect, "/#")
	}
}
