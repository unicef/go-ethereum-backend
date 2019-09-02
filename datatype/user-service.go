package datatype

import "github.com/qjouda/dignity-platform/backend/decimaldt"

//UserService defines asset service interface
type UserService interface {
	Register(
		email string,
		password string,
		username string,
		ethereumAdderss string,
		agreeToTerms bool,
	) (*User, error)
	Validate(string, string) error
	FindByID(ID) (*User, error)
	FindByEmail(string) (*User, error)
	IsEmailRegistered(string) bool
	FindUserBalances(userID ID) ([]Balance, error)
	FindUserBalance(userID ID, assetID ID) (
		availableBalance decimaldt.Decimal,
		reserved decimaldt.Decimal,
		err error,
	)
	NewPassResetTokenByEmail(email string) (*UserResetToken, error)
	ConfirmPassResetToken(userID ID, token, newPassword string) error
	ChangePassword(*User, string, string) error
	CreateChangeEmailConfirmation(*User, string, string) (*EmailChangeConfirmation, error)
	ConfirmUserEmailChange(ID, string) bool
	UpdateProfileImage(user *User, imageURL string) error
	FindUserNotifications(id ID, offset int, limit int) ([]NotificationMsg, error)
	Notify(
		senderID ID,
		receiverID ID,
		topicID ID,
		notiType NotificationType,
	) error
	MarkAsRead(ID ID) error
	MarkAllAsRead(user *User) error
	FindEmojis() ([]Emoji, error)
	AddUserPostReaction(
		user *User,
		blockID ID,
		reactionID ID,
	) error
}
