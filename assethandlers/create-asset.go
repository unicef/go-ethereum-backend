package assethandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qjouda/dignity-platform/backend/auth"
	"github.com/qjouda/dignity-platform/backend/datatype"
)

// This is a general interface for Currency tokens passed from front-ends
// The backend should enable initition of these and deployment on the used blockchain
// AssetProposalBody Asset body type: used to decrypt asset data passed from API users
// type AssetProposalBody struct {
// 	Name    string `json:"name"`
// 	Purpose string `json:"purpose"`
// 	Symbol  string `json:"symbol"`
// 	Supply  struct {
// 		// allocate unconditional initial supply to the provided addresses
// 		Initial map[string]float64 `json:"initial"`
// 		// Cap: pass int value to cap the asset or 0 for uncapped
// 		Cap int `json:"cap"`
// 	} `json:"Supply"`
// 	// asset units allocations, ethereum address to percentage (0.0 - 1.0)
// 	Allocation map[string]float32 `json:"Allocation"`
// }

// CreateAsset creates a new asset proposal
func CreateAsset(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := auth.MustGetUser(c)
		body := struct {
			Name    string `json:"name"`
			Purpose string `json:"purpose"`
			Symbol  string `json:"symbol"`
		}{}
		c.BindJSON(&body)
		market, err := sc.AssetService.Create(
			user.ID,
			body.Name,
			body.Symbol,
			body.Purpose,
		)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, market)
	}
}
