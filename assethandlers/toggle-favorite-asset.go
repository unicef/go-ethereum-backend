package assethandlers

import (
	"net/http"

	"github.com/qjouda/dignity-platform/backend/auth"
	"github.com/qjouda/dignity-platform/backend/datatype"
	"github.com/gin-gonic/gin"
)

//ToggleFavoriteAsset toggles favorite asset
func ToggleFavoriteAsset(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {
		u := auth.MustGetUser(c)
		body := struct {
			AssetID int
		}{}
		c.BindJSON(&body)
		err := sc.AssetService.ToggleFavorite(sc, u, datatype.ID(body.AssetID))
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, nil)
	}
}
