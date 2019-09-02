package assetservice

import (
	"errors"

	"github.com/qjouda/dignity-platform/backend/apperrors"
	"github.com/qjouda/dignity-platform/backend/datatype"
	"github.com/qjouda/dignity-platform/backend/decimaldt"
)

// Create insert asset
func (db *Service) Transfer(
	sc datatype.ServiceContainer,
	fromID datatype.ID,
	toID datatype.ID,
	amount decimaldt.Decimal,
	assetID datatype.ID,
) error {
	balance, _, err := sc.UserService.FindUserBalance(fromID, assetID)
	if err != nil {
		return err
	}
	if balance.IsLessThan(amount) {
		return errors.New("Not enough balance")
	}
	// Start db transaction
	tx, err := db.Begin()
	if err != nil {
		apperrors.Critical("assetservice:transfer:1", err)
		return err
	}
	defer tx.Rollback()
	/**** Start transaction ***/
	// Update from balance
	_, err = tx.Exec(
		`UPDATE user_balance SET
					balance = ?
    WHERE assetId = ? AND userId = ?`,
		balance.Sub(amount),
		datatype.ID(assetID),
		fromID,
	)
	if err != nil {
		apperrors.Critical("assetservice:transfer:2", err)
		return datatype.ErrServerError
	}
	toBalance, _, err := sc.UserService.FindUserBalance(toID, assetID)
	if err != nil {
		return err
	}
	_, err = tx.Exec(
		`UPDATE user_balance SET
					balance = ?
    WHERE assetId = ? AND userId = ?`,
		toBalance.Add(amount),
		datatype.ID(assetID),
		toID,
	)
	if err != nil {
		apperrors.Critical("assetservice:transfer:3", err)
		return datatype.ErrServerError
	}
	/*** Commit db transaction ***/
	err = tx.Commit()
	if err != nil {
		apperrors.Critical("assetservice:transfer:4", err)
		return datatype.ErrServerError
	}
	return nil
}
