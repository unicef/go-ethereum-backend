package datatype

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User user type
type User struct {
	ID              ID        `json:"id"`
	Email           string    `json:"email"`
	Name            string    `json:"name"`
	ProfileImageURL string    `json:"profileImageURL"`
	EthereumAddress string    `json:"ethereumAddress"`
	Password        string    `json:"-"`
	Salt            string    `json:"-"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
	AgreeToTerms    DbBool    `json:"agreeToTerms"`
	IsDeleted       DbBool    `json:"-"`
}

// NewUser creates a new user
func NewUser(randomstring RandomString) *User {
	return &User{
		Salt:         randomstring.Generate(32),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		AgreeToTerms: DbFalse,
		IsDeleted:    DbFalse,
	}
}

// IsPassword compares given unencrypted-password to user password
func (user *User) IsPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password+user.Salt))
	return err == nil
}
