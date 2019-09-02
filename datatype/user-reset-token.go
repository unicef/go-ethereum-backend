package datatype

import "time"

// UserResetToken reset token struct
type UserResetToken struct {
	ID        ID
	UserID    ID
	Token     string
	CreatedAt time.Time
}

// NewUserResetToken creates a new reset password token
func NewUserResetToken(userID ID, randomstring RandomString) *UserResetToken {
	return &UserResetToken{
		UserID:    userID,
		Token:     randomstring.Generate(64),
		CreatedAt: time.Now(),
	}
}
