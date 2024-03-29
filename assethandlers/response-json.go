package assethandlers

import (
	"github.com/qjouda/dignity-platform/backend/datatype"
)

func toAssetsResponse(entries []datatype.Asset) []interface{} {
	res := []interface{}{}
	for _, entry := range entries {
		res = append(res, struct {
			ID               datatype.ID
			Symbol           string
			Name             string
			CreatorID        datatype.ID
			CreatorName      string
			Description      string
			Supply           int64
			MinersCounter    int
			FavoritesCounter int
			DidUserLike      bool
		}{
			ID:               entry.ID,
			Symbol:           entry.Symbol,
			Name:             entry.Name,
			CreatorID:        entry.CreatorID,
			CreatorName:      entry.CreatorName,
			Description:      entry.Description,
			Supply:           entry.Supply,
			MinersCounter:    entry.MinersCounter,
			FavoritesCounter: entry.FavoritesCounter,
			DidUserLike:      entry.DidUserLike,
		})
	}
	return res
}
