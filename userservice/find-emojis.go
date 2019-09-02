package userservice

import (
	"github.com/qjouda/dignity-platform/backend/apperrors"
	"github.com/qjouda/dignity-platform/backend/datatype"
)

// FindEmojis finds avaialble emoji reactions
func (db *Service) FindEmojis() ([]datatype.Emoji, error) {
	result := []datatype.Emoji{}
	rows, err := db.Query(`
		SELECT
			id,
			name,
			logo
		FROM reactions`)
	if err != nil {
		apperrors.Critical("user:userservice:find-emojis:1", err)
		return result, datatype.ErrServerError
	}
	defer rows.Close()
	for rows.Next() {
		entry := datatype.Emoji{}
		err = rows.Scan(
			&entry.ID,
			&entry.Name,
			&entry.Logo,
		)
		if err != nil {
			apperrors.Critical("user:userservice:find-emojis:2", err)
			return result, datatype.ErrServerError
		}
		result = append(result, entry)
	}
	if err = rows.Err(); err != nil {
		apperrors.Critical("user:userservice:find-emojis:3", err)
		return result, datatype.ErrServerError
	}
	return result, nil
}
