package assethandlers

import (
	"net/http"

	"github.com/qjouda/dignity-platform/backend/auth"
	"github.com/qjouda/dignity-platform/backend/datatype"
	"github.com/qjouda/dignity-platform/backend/decimaldt"
	"github.com/gin-gonic/gin"
)

// Transfer creat a new asset
func Transfer(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := auth.MustGetUser(c)
		body := struct {
			Amount  decimaldt.Decimal `json:"amount"`
			ToID    datatype.ID       `json:"toId"`
			AssetID datatype.ID       `json:"assetId"`
		}{}
		c.BindJSON(&body)
		err := sc.AssetService.Transfer(
			sc,
			user.ID,
			body.ToID,
			body.Amount,
			body.AssetID,
		)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, nil)
	}
}
