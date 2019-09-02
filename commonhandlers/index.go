package commonhandlers

import (
	"github.com/gin-gonic/gin"
	"github.com/qjouda/dignity-platform/backend/datatype"
)

// Index index route
func Index(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.File("public/index.html")
	}
}
