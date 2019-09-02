package userservice

import (
	"github.com/qjouda/dignity-platform/backend/apperrors"
	"github.com/qjouda/dignity-platform/backend/datatype"
)

// MarkAllAsRead mark all notifications as read for the passed user
func (db *Service) MarkAllAsRead(
	user *datatype.User,
) error {
	_, err := db.Exec(`
    UPDATE not_message SET
      isRead = 1
    WHERE receiverId = ?`,
		user.ID,
	)
	if err != nil {
		apperrors.Critical("userservice:MarkAllAsRead:1", err)
		return err
	}
	return nil
}
