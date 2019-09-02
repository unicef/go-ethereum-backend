package userhandlers

import (
	"net/http"

	"github.com/qjouda/dignity-platform/backend/auth"
	"github.com/qjouda/dignity-platform/backend/datatype"
	"github.com/qjouda/dignity-platform/backend/decimaldt"
	"github.com/gin-gonic/gin"
)

// GetUserBalances gets portfolio position similar to UserGetBalances,
// but with extra info
func GetUserBalances(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := auth.MustGetUser(c)
		balances, err := sc.UserService.FindUserBalances(user.ID)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, toBalancesResponse(balances))
	}
}

func toBalancesResponse(entries []datatype.Balance) map[string]interface{} {
	positions := []interface{}{}
	for _, entry := range entries {
		balance := entry.Balance.Add(entry.Reserved)
		if balance.Equal(decimaldt.NewFromFloat(0)) {
			continue
		}
		positions = append(positions, struct {
			AssetID  datatype.ID
			Symbol   string
			Name     string
			Reserved string
			Balance  string
		}{
			AssetID:  entry.AssetID,
			Symbol:   entry.AssetSymbol,
			Name:     entry.AssetName,
			Reserved: entry.Reserved.String(),
			Balance:  balance.String(),
		})
	}
	res := make(map[string]interface{})
	res["Balances"] = positions
	return res
}
