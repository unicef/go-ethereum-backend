package datatype

import "github.com/qjouda/dignity-platform/backend/decimaldt"

//AssetService defines asset service interface
type AssetService interface {
	FindAll(user *User) ([]Asset, error)
	FindByID(id ID) (*Asset, error)
	FindBySymbol(symbol string) (*Asset, error)
	FindByName(name string) (*Asset, error)
	Create(
		userID ID,
		name string,
		symbol string,
		description string,
	) (*Asset, error)
	ToggleFavorite(sc ServiceContainer, user *User, assetID ID) error
	ToggleFavoriteBlock(user *User, blockID ID) error
	FindUserFavoriteAssets(user *User) ([]Asset, error)
	CreateAssetBlock(
		userID ID,
		assetID ID,
		blockText string,
		images []string,
	) (*Block, error)
	GetAssetBlocks(assetID ID) ([]Block, error)
	VerifyAssetBlock(sc ServiceContainer, user *User, blockID ID, status int) error
	GetAssetMiners(assetID ID) ([]Miner, error)
	IsOracle(userID ID, assetID ID) bool
	GetAssetBlockImages(blockID ID) ([]string, error)
	Transfer(
		sc ServiceContainer,
		fromID ID,
		toID ID,
		amount decimaldt.Decimal,
		assetID ID,
	) error
	DidUserLike(userID ID, assetID ID) (bool, error)
}
