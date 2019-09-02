package assethandlers

import (
	"net/http"

	"github.com/qjouda/dignity-platform/backend/auth"
	"github.com/qjouda/dignity-platform/backend/datatype"
	"github.com/gin-gonic/gin"
)

// GetAssets fetch assets from db
func GetAssets(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := auth.MustGetUser(c)
		assets, err := sc.AssetService.FindAll(user)
		if err != nil {
			c.String(http.StatusServiceUnavailable, err.Error())
			return
		}
		c.JSON(http.StatusOK, toAssetsResponse(assets))
	}
}
