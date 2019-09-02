package assetservice

import (
	"database/sql"

	"github.com/qjouda/dignity-platform/backend/apperrors"
	"github.com/qjouda/dignity-platform/backend/datatype"
)

//IsMiner checks whether a user is a miner of a token or not
func (db *Service) IsMiner(userID datatype.ID, assetID datatype.ID) bool {
	var id datatype.ID
	err := db.QueryRow(
		`SELECT
			assetId as counter
		FROM asset_block
		WHERE assetId=? AND userId=? and status=? LIMIT 1`,
		assetID,
		userID,
		datatype.BlockAccepted,
	).Scan(&id)
	if err == sql.ErrNoRows {
		return false
	}
	if err != nil {
		apperrors.Critical("assetservice:IsMiner:1", err)
		return false
	}
	return true
}
