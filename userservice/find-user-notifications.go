package userservice

import (
	"github.com/qjouda/dignity-platform/backend/apperrors"
	"github.com/qjouda/dignity-platform/backend/datatype"
	"github.com/qjouda/dignity-platform/backend/helpers"
)

// FindUserNotifications ...
func (db *Service) FindUserNotifications(
	id datatype.ID,
	offset int,
	limit int,
) ([]datatype.NotificationMsg, error) {
	result := []datatype.NotificationMsg{}
	rows, err := db.Query(`
    SELECT
			noti.id,
			noti.senderId,
      sender.username,
      noti.receiverId,
      receiver.username,
			noti.topicId,
      asset.name,
      asset.symbol,
      noti.type,
      noti.isRead,
      noti.createdAt
		FROM not_message noti
    LEFT JOIN user sender ON sender.id=noti.senderId
    LEFT JOIN user receiver ON receiver.id=noti.receiverId
    LEFT JOIN asset asset ON noti.topicId=asset.Id
		WHERE noti.receiverId=?
		ORDER BY noti.createdAt DESC`,
		id,
	)
	if err != nil {
		apperrors.Critical("userservices:fin-user-notification:1", err)
		return result, datatype.ErrServerError
	}
	defer rows.Close()
	for rows.Next() {
		entry := datatype.NotificationMsg{}
		err = rows.Scan(
			&entry.ID,
			&entry.SenderID,
			&entry.SenderName,
			&entry.ReceiverID,
			&entry.ReceiverName,
			&entry.TopicID,
			&entry.TopicName,
			&entry.TopicSymbol,
			&entry.Type,
			&entry.IsRead,
			&entry.CreatedAt,
		)
		if err != nil {
			apperrors.Critical("userservice:fin-user-notification:2", err)
			return result, datatype.ErrServerError
		}
		entry.CreatedAtHuman = helpers.DateToHuman(entry.CreatedAt)
		result = append(result, entry)
	}
	if err = rows.Err(); err != nil {
		apperrors.Critical("userservice:fin-user-notification:3", err)
		return result, datatype.ErrServerError
	}
	return result, nil
}
