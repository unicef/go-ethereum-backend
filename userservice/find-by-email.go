package userservice

import (
	"fmt"

	"github.com/qjouda/dignity-platform/backend/datatype"
)

//FindByEmail finds a user by email
func (db *Service) FindByEmail(email string) (*datatype.User, error) {
	row := db.QueryRow(
		fmt.Sprintf("SELECT %s FROM user WHERE email=?", getUserCols()),
		email,
	)

	return scanUser(row)
}
