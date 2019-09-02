package assetservice

import "github.com/qjouda/dignity-platform/backend/dbservice"

//Service defines asset service type
type Service struct {
	*dbservice.DB
}

//NewService factory for Service
func NewService(db *dbservice.DB) *Service {
	return &Service{db}
}

func getAssetCols() string {
	return `
		asset.id,
		asset.symbol,
		asset.name,
		asset.creatorId,
		asset.description,
		asset.supply,
		asset.decimals,
		asset.minersCounter,
		asset.favoritesCount`
}
