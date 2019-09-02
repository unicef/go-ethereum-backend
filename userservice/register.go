package userservice

import (
	"errors"
	"strings"

	"github.com/qjouda/dignity-platform/backend/appstrings"
	"github.com/qjouda/dignity-platform/backend/datatype"
	"github.com/lytics/logrus"
)

//Register Registers a new user
func (db *Service) Register(
	email,
	password string,
	username string,
	ethereumAddress string,
	agreeToTerms bool,
) (*datatype.User, error) {
	var err error
	email = sanitizeEmail(email)
	password = strings.TrimSpace(password)
	if err = db.Validate(email, password); err != nil {
		return nil, err
	}
	if len(username) > 64 || len(username) < 2 {
		return nil, errors.New("User name has to be at least 2 characters and less than 64")
	}
	if db.IsUsernameRegistered(username) {
		return nil, datatype.ErrUsernameExists
	}
	if !agreeToTerms {
		return nil, datatype.ErrAgreeToTerms
	}
	user := datatype.NewUser(appstrings.NewRandom())
	user.Email = email
	user.Password, err = hashPassword(password + user.Salt)
	user.Name = username
	if err != nil {
		logrus.Error("Regisyer:1")
		return nil, datatype.ErrServerError
	}

	res, err := db.Exec(
		`INSERT INTO user SET
			email = ?,
			username = ?,
			password = ?,
			salt = ?,
			createdAt = ?,
			updatedAt = ?,
			ethereumAddress=?,
			agreeToTerms = ?,
			isDeleted = ?`,
		user.Email,
		user.Name,
		user.Password,
		user.Salt,
		user.CreatedAt,
		user.UpdatedAt,
		ethereumAddress,
		user.AgreeToTerms,
		user.IsDeleted,
	)
	if err != nil {
		logrus.WithFields(
			logrus.Fields{"e": err.Error()},
		).Error("userservice:register:1")
		return nil, datatype.ErrServerError
	}
	id, err := res.LastInsertId()
	if err != nil {
		logrus.WithFields(
			logrus.Fields{"e": err.Error()},
		).Error("userservice:register:2")
		return nil, datatype.ErrServerError
	}
	return db.FindByID(datatype.ID(id))
}
