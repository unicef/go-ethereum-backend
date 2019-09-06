package userhandlers

import (
	"github.com/gin-gonic/gin"
	"github.com/qjouda/dignity-platform/backend/datatype"
	"github.com/qjouda/dignity-platform/backend/routermiddleware"
)

// InjectHandlers injects user handlers into the application main router [dependency injection]
func InjectHandlers(sc datatype.ServiceContainer, rg *gin.RouterGroup) {
	authenticator := routermiddleware.SessionMustAuth()
	rg.POST("/register", Register(sc))
	rg.POST("/login", Login(sc))
	rg.GET("/logout", Logout(sc))
	rg.GET("/session", authenticator, SessionGet(sc))
	rg.DELETE("/session", authenticator, SessionDestroy(sc))
	rg.POST("/forgotpass-requests/new", PasswordRequestNew(sc))
	rg.POST("/forgotpass-requests/reset", PasswordRequestReset(sc))
	rg.POST("/user/passwords", authenticator, ChangePassword(sc))
	rg.POST("/user/email", authenticator, ChangeEmail(sc))
	rg.GET("/user/email/confirm", ConfirmChangeEmail(sc))
	rg.GET("/balances", authenticator, GetUserBalances(sc))
	rg.GET("/notifications", authenticator, FindUserNotifications(sc))
	rg.POST("/notification/mark-as-read", authenticator, MarkAsRead(sc))
	rg.POST("/notification/mark-all-as-read", authenticator, MarkAllAsRead(sc))
	rg.GET("/users/:userID", authenticator, GetUser(sc))
	rg.POST("/user/profile-image", authenticator, UploadProfileImage(sc))
	rg.GET("/emojis", ListEmojis(sc))
	rg.POST("/user/submit-emoji", authenticator, PostReaction(sc))
}
