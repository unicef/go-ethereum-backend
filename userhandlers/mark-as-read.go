package userhandlers

import (
	"net/http"

	"github.com/qjouda/dignity-platform/backend/datatype"
	"github.com/gin-gonic/gin"
)

// MarkAsRead mark as read
func MarkAsRead(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {
		body := struct {
			ID datatype.ID `json:"id"`
		}{}
		c.BindJSON(&body)
		err := sc.UserService.MarkAsRead(
			body.ID,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, nil)
	}
}
