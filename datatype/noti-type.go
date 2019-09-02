package datatype

import "time"

// NotificationType a data type representing notification type (NewFollow, NewPost, PostAccepted)
type NotificationType uint8

const (
	// NEWFOLLOW notification type: new follow noti
	NEWFOLLOW = 0
	// NEWPOST notification type: new post
	NEWPOST = 1
	// POSTACCEPTED notification type:  post accepted
	POSTACCEPTED = 2
	// POSTLIKE notification type: new post like
	POSTLIKE = 3
)

// NotificationMsg notification message data type
type NotificationMsg struct {
	ID             ID
	SenderID       ID
	SenderName     string
	ReceiverID     ID
	ReceiverName   string
	TopicID        ID
	TopicName      string
	TopicSymbol    string
	Type           NotificationType
	IsRead         bool
	CreatedAt      time.Time
	CreatedAtHuman string
}
