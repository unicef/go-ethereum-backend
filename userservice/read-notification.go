package userservice

import (
	"github.com/qjouda/dignity-platform/backend/apperrors"
	"github.com/qjouda/dignity-platform/backend/datatype"
)

// MarkAsRead mark a notiffication message as read
func (db *Service) MarkAsRead(
	ID datatype.ID,
) error {
	_, err := db.Exec(`
    UPDATE not_message SET
      isRead = 1
    WHERE id = ?`,
		ID,
	)
	if err != nil {
		apperrors.Critical("userservice:MarkAsRead:1", err)
		return err
	}
	return nil
}
