package assethandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qjouda/dignity-platform/backend/auth"
	"github.com/qjouda/dignity-platform/backend/datatype"
)

//GetUserFavoriteAssets get user's favorite assets
func GetUserFavoriteAssets(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {
		u := auth.MustGetUser(c)
		assets, err := sc.AssetService.FindUserFavoriteAssets(u)
		if err != nil {
			c.String(http.StatusServiceUnavailable, err.Error())
			return
		}
		c.JSON(http.StatusOK, toAssetsResponse(assets))
	}
}
