package assetservice

import (
	"github.com/qjouda/dignity-platform/backend/apperrors"
	"github.com/qjouda/dignity-platform/backend/datatype"
)

// GetAssetBlockImages finds block images
func (db *Service) GetAssetBlockImages(
	blockID datatype.ID,
) ([]string, error) {
	result := []string{}
	rows, err := db.Query(
		`SELECT
      filepath
    FROM asset_block_image
    WHERE blockId = ?`,
		blockID,
	)
	if err != nil {
		apperrors.Critical("assetservice:GetAssetBlockImages:1", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var image string
		err := rows.Scan(
			&image,
		)
		if err != nil {
			apperrors.Critical("assetservice:GetAssetBlockImages:2", err)
			return nil, err
		}
		result = append(result, image)
	}
	if err := rows.Err(); err != nil {
		apperrors.Critical("assetservice:GetAssetBlockImages:3", err)
		return nil, err
	}
	return result, nil
}
