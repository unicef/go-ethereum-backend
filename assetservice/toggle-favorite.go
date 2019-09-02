package assetservice

import (
	"errors"
	"time"

	"github.com/qjouda/dignity-platform/backend/apperrors"
	"github.com/qjouda/dignity-platform/backend/datatype"
)

//ToggleFavorite toggles asset as fav/not fav
func (db *Service) ToggleFavorite(
	sc datatype.ServiceContainer,
	user *datatype.User,
	assetID datatype.ID,
) error {
	if user == nil {
		return errors.New("User is missing")
	}
	count := 0
	err := db.QueryRow(
		"SELECT count(*) FROM asset WHERE id = ?",
		assetID,
	).Scan(&count)
	if err != nil {
		return datatype.ErrServerError
	}
	if count != 1 {
		return errors.New("Invalid asset")
	}
	err = db.QueryRow(
		"SELECT count(*) FROM asset_favorites WHERE userId = ? AND assetId = ?",
		user.ID,
		assetID,
	).Scan(&count)
	if err != nil {
		apperrors.Critical("assetservice:toggle-favorite-block:1", err)
		return datatype.ErrServerError
	}
	if count > 0 {
		_, err := db.Exec(
			"DELETE FROM asset_favorites WHERE userId = ? AND assetId = ?",
			user.ID,
			assetID,
		)
		if err != nil {
			apperrors.Critical("assetservice:toggle-favorite-block:2", err)
			return datatype.ErrServerError
		}
		_, err = db.Exec(`UPDATE asset
			SET favoritesCount = favoritesCount - 1
			WHERE id = ? `,
			assetID,
		)
		if err != nil {
			apperrors.Critical("assetservice:toggle-favorite-block:3", err)
			return datatype.ErrServerError
		}
	} else {
		_, err := db.Exec(
			"INSERT INTO asset_favorites SET userId = ?, assetId = ?, addedAt = ?",
			user.ID,
			assetID,
			time.Now(),
		)
		if err != nil {
			apperrors.Critical("assetservice:toggle-favorite-block:4", err)
			return datatype.ErrServerError
		}
		_, err = db.Exec(`UPDATE asset
			SET favoritesCount = favoritesCount + 1
			WHERE id = ? `,
			assetID,
		)
		if err != nil {
			apperrors.Critical("assetservice:toggle-favorite-block:5", err)
			return datatype.ErrServerError
		}
		asset, _ := sc.AssetService.FindByID(assetID)
		sc.UserService.Notify(user.ID, asset.CreatorID, assetID, datatype.NEWFOLLOW)
	}
	return nil
}
