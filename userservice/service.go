package userservice

import (
	"github.com/asaskevich/govalidator"
	"github.com/qjouda/dignity-platform/backend/datatype"
	"github.com/qjouda/dignity-platform/backend/dbservice"
)

//Service defines service type
type Service struct {
	*dbservice.DB
}

//NewService factory for Service
func NewService(db *dbservice.DB) *Service {
	return &Service{db}
}

// Validate validates email and password for signup
func (db *Service) Validate(email string, password string) error {
	if !govalidator.IsEmail(email) {
		return datatype.ErrEmailInvalid
	}
	if db.IsEmailRegistered(email) {
		return datatype.ErrEmailExists
	}
	if err := validatePassword(password, email); err != nil {
		return err
	}
	return nil
}
