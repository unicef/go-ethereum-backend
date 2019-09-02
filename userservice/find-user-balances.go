package userservice

import (
	"github.com/qjouda/dignity-platform/backend/datatype"
	"github.com/sirupsen/logrus"
)

// FindUserBalances finds all balances belonging to given user
func (db *Service) FindUserBalances(id datatype.ID) ([]datatype.Balance, error) {
	result := []datatype.Balance{}
	rows, err := db.Query(`SELECT
			b.assetId,
			b.balance,
			b.reserved,
			c.name,
			c.symbol
		FROM user_balance b
		LEFT JOIN
			asset c ON c.id = b.assetId
		WHERE b.userId=?`,
		id,
	)
	if err != nil {
		logrus.WithFields(
			logrus.Fields{"e": err.Error(), "userID": id},
		).Error("user:FindUserBalances:1")
		return result, datatype.ErrServerError
	}
	defer rows.Close()
	for rows.Next() {
		entry := datatype.Balance{UserID: id}
		err = rows.Scan(
			&entry.AssetID,
			&entry.Balance,
			&entry.Reserved,
			&entry.AssetName,
			&entry.AssetSymbol,
		)
		if err != nil {
			logrus.WithFields(
				logrus.Fields{"e": err.Error(), "userID": id},
			).Error("user:FindUserBalances:2")
			return result, datatype.ErrServerError
		}
		result = append(result, entry)
	}
	if err = rows.Err(); err != nil {
		logrus.WithFields(
			logrus.Fields{"e": err.Error(), "userID": id},
		).Error("user:FindUserBalances:3")
		return result, datatype.ErrServerError
	}
	return result, nil
}
