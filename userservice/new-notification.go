package userservice

import (
	"time"

	"github.com/qjouda/dignity-platform/backend/apperrors"
	"github.com/qjouda/dignity-platform/backend/datatype"
)

// Notify creates a new notification  entry
func (db *Service) Notify(
	senderID datatype.ID,
	receiverID datatype.ID,
	topicID datatype.ID,
	notiType datatype.NotificationType,
) error {
	if senderID == receiverID {
		return nil
	}
	_, err := db.Exec(`
    INSERT INTO not_message SET
      senderId = ?,
      receiverId = ?,
      topicId = ?,
      type = ?,
      isRead = 0,
      createdAt = ?`,
		senderID,
		receiverID,
		topicID,
		notiType,
		time.Now(),
	)
	if err != nil {
		apperrors.Critical("userservice:new-notification:1", err)
		return err
	}
	return nil
}
