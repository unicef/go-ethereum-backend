package assetservice

import (
	"errors"
	"strings"

	"github.com/qjouda/dignity-platform/backend/apperrors"
	"github.com/qjouda/dignity-platform/backend/datatype"
)

//Create insert asset
func (db *Service) Create(
	userID datatype.ID,
	name string,
	symbol string,
	description string,
) (*datatype.Asset, error) {
	name = strings.TrimSpace(name)
	symbol = strings.TrimSpace(symbol)
	description = strings.TrimSpace(description)
	e := db.ValidateAsset(name, symbol, description)
	if e != nil {
		return nil, e
	}
	{ // check if asset name and symbol dont exist already
		a, err := db.FindByName(name)
		if err != nil {
			apperrors.Critical("assetservice:create:1", err)
			return nil, datatype.ErrServerError
		}
		if a != nil {
			return nil, errors.New("Asset with (" + name + ") name already exists. Please try a different name")
		}
		a, err = db.FindBySymbol(symbol)
		if err != nil {
			apperrors.Critical("assetservice:create:2", err)
			return nil, datatype.ErrServerError
		}
		if a != nil {
			return nil, errors.New("Asset with (" + symbol + ") symbol already exists. Please try a different symbol")
		}
	}
	// Start db transaction
	tx, err := db.Begin()
	if err != nil {
		apperrors.Critical("assetservice:create:3", err)
		return nil, err
	}
	defer tx.Rollback()
	/**** Start transaction ***/
	// Insert Asset, market, and increase balance of user on one run or fail
	// insert in asset table
	res, err := tx.Exec(
		`INSERT INTO asset SET
	        name = ?,
					symbol = ?,
					description = ?,
					supply = 0,
					creatorId = ?,
					decimals = 8
	      `,
		name,
		symbol,
		description,
		userID,
	)
	if err != nil {
		apperrors.Critical("assetservice:create:4", err)
		return nil, datatype.ErrServerError
	}
	assetID, err := res.LastInsertId()
	if err != nil {
		apperrors.Critical("assetservice:create:5", err)
		return nil, datatype.ErrServerError
	}
	// Update user balance
	res, err = tx.Exec(
		`INSERT INTO user_balance SET
					userId = ?,
					assetID = ?,
					balance = ?,
					reserved = 0
	      `,
		userID,
		datatype.ID(assetID),
		"0",
	)
	if err != nil {
		apperrors.Critical("assetservice:create:8", err)
		return nil, datatype.ErrServerError
	}
	/*** Commit db transaction ***/
	err = tx.Commit()
	if err != nil {
		apperrors.Critical("assetservice:create:9", err)
		return nil, datatype.ErrServerError
	}
	market, err := db.FindByID(datatype.ID(assetID))
	if err != nil {
		apperrors.Critical("assetservice:create:10", err)
		return nil, datatype.ErrServerError
	}
	return market, nil
}
