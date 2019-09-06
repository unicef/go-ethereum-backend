package userhandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qjouda/dignity-platform/backend/datatype"
)

// ListEmojis handler for available empojis
func ListEmojis(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {
		emojis, err := sc.UserService.FindEmojis()
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, struct {
			Entries []datatype.Emoji
		}{
			Entries: emojis,
		})
	}
}
