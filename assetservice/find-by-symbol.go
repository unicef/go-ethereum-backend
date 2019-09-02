package assetservice

import (
	"database/sql"

	"github.com/qjouda/dignity-platform/backend/datatype"
	"github.com/qjouda/dignity-platform/backend/helpers"
)

// FindBySymbol find asset by symbol
func (db *Service) FindBySymbol(symbol string) (*datatype.Asset, error) {
	var c datatype.Asset
	err := db.QueryRow(`
		SELECT
			asset.id,
			asset.name,
			asset.symbol,
			asset.description,
			asset.supply,
			asset.creatorId,
			user.username,
			asset.minersCounter,
			asset.favoritesCount,
			asset.ethereumAddress,
			asset.ethereumTransactionAddress,
			asset.createdAt
		FROM
			asset asset
		LEFT JOIN user user ON asset.creatorId=user.id
		WHERE asset.symbol = ?`,
		symbol,
	).Scan(
		&c.ID,
		&c.Name,
		&c.Symbol,
		&c.Description,
		&c.Supply,
		&c.CreatorID,
		&c.CreatorName,
		&c.MinersCounter,
		&c.FavoritesCounter,
		&c.EthereumAddress,
		&c.EthereumTransactionAddress,
		&c.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	c.CreatedAtHuman = helpers.DateToHuman(c.CreatedAt)
	return &c, err
}
