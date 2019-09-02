package ethereum

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/lytics/logrus"
	"github.com/qjouda/dignity-platform/backend/ethereum/generated"
)

// Mint mints a new currency units to the passed token and toAddress
func (b *Ethereum) Mint(
	tokenAddress common.Address,
	toAddress common.Address,
	amount int64,
) (*types.Transaction, error) {
	mintable, err := generated.NewMintable(tokenAddress, b.rpc)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("ethereum:Mint:1")
		return nil, err
	}
	// @TODO change b.auth to the address of the user who is minting the new
	// tokens i.e claim approver
	// This is now abstraced in the server wallet for easiness, however,
	// this should be passed from the front-end when user has their own wallet
	// applications
	txAddress, err := mintable.Mint(b.auth, toAddress, big.NewInt(amount))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("ethereum:Mint:2")
		return nil, err
	}
	return txAddress, nil
}
