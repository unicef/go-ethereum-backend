package datatype

import (
	"github.com/qjouda/dignity-platform/backend/decimaldt"
)

// Miner user type
type Miner struct {
	ID               ID
	UserName         string
	Mined            decimaldt.Decimal
	MiningPercentage string
}
