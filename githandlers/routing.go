package githandlers

import (
	"github.com/gin-gonic/gin"
	"github.com/qjouda/dignity-platform/backend/datatype"
	"github.com/qjouda/dignity-platform/backend/routermiddleware"
)

//InjectHandlers injects asset handlers into the application
func InjectHandlers(sc datatype.ServiceContainer, rg *gin.RouterGroup) {
	authenticator := routermiddleware.SessionMustAuth()
	rg.POST("/load-repo-details", authenticator, LoadRepoDetails(sc))
}
