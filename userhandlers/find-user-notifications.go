package userhandlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/qjouda/dignity-platform/backend/auth"
	"github.com/qjouda/dignity-platform/backend/datatype"
)

// FindUserNotifications finds user related notifications
func FindUserNotifications(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := auth.MustGetUser(c)
		page, err := strconv.Atoi(c.Query("page"))
		if err != nil {
			page = 1
		}
		limit := 25
		offset := page*limit - limit

		entries, err := sc.UserService.FindUserNotifications(
			user.ID,
			offset,
			limit,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, entries)
	}
}
