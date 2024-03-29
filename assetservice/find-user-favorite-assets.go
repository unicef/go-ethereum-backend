package assetservice

import (
	"github.com/qjouda/dignity-platform/backend/datatype"
	"github.com/qjouda/dignity-platform/backend/apperrors"
	"github.com/qjouda/dignity-platform/backend/helpers"
)

//FindUserFavoriteAssets finds all user's favored assits
func (db *Service) FindUserFavoriteAssets(user *datatype.User) ([]datatype.Asset, error) {
	result := []datatype.Asset{}
	rows, err := db.Query(`
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
			asset.createdAt,
			IF(favorites.blockId, TRUE, FALSE),
		FROM asset asset
		LEFT JOIN user user ON asset.creatorId = user.id
		LEFT JOIN asset_favorites favorites ON asset.id=favorites.assetId AND favorites.userId=?
		ORDER BY id DESC
		WHERE asset.id IN (
			SELECT assetId from asset_favorites WHERE userId = ? ORDER BY createdAt DESC
		)`,
		user.ID,
		user.ID,
	)
	if err != nil {
		apperrors.Critical("assetservice:find-user-favorite-assets:0", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var c datatype.Asset
		err := rows.Scan(
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
			&c.CreatedAtHuman,
			&c.DidUserLike,
		)
		if err != nil {
			apperrors.Critical("assetservice:find-user-favorite-assets:1", err)
			return nil, err
		}
		c.CreatedAtHuman = helpers.DateToHuman(c.CreatedAt)
		result = append(result, c)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}
