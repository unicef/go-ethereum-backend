package auth

import (
	"fmt"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/qjouda/dignity-platform/backend/datatype"
)

// Login logs user in
func Login(c *gin.Context, user *datatype.User) {
	session := sessions.Default(c)
	session.Set("userId", fmt.Sprintf("%d", user.ID))
	session.Save()
}

// Logout logs user out
func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("userId")
	session.Save()
}

//MustGetUser gets logged in user or panic
func MustGetUser(c *gin.Context) *datatype.User {
	return c.MustGet("user").(*datatype.User)
}
