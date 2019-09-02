package ethereum

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/qjouda/dignity-platform/backend/ethereum/generated"
)

// DeployMintable deployes new Mintable token to Ethereum from the server
// account returns the newly deployed token address and the corresponding
// transaction
func (b *Ethereum) DeployMintable(
	name_ string,
	symbol_ string,
	decimals_ uint8,
	minter common.Address,
) (common.Address, *types.Transaction, error) {
	address, tx, _, err := generated.DeployMintable(
		b.auth,
		b.rpc,
		name_,
		symbol_,
		decimals_,
		minter,
	)
	if err != nil {
		return address, nil, err
	}
	return address, tx, nil
}
