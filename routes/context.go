package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/qjouda/dignity-platform/backend/datatype"
)

func mustGetUser(c *gin.Context) *datatype.User {
	return c.MustGet("user").(*datatype.User)
}
