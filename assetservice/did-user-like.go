package assetservice

import (
	"github.com/qjouda/dignity-platform/backend/datatype"
)

// DidUserLike di user like
func (db *Service) DidUserLike(userID datatype.ID, assetID datatype.ID) (bool, error) {
	didUserLike := false
	err := db.QueryRow(
		`SELECT
			IF(favorites.assetId, TRUE, FALSE)
		FROM asset
		LEFT JOIN asset_favorites favorites ON asset.id=favorites.assetId AND favorites.userId=?
    WHERE asset.id=?`,
		userID,
		assetID,
	).Scan(&didUserLike)
	if err != nil {
		return false, err
	}
	return didUserLike, nil
}
