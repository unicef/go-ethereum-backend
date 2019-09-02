package datatype

import (
	"github.com/qjouda/dignity-platform/backend/decimaldt"
)

// Balance balance type
type Balance struct {
	UserID      ID
	AssetID     ID
	AssetName   string
	AssetSymbol string
	Balance     decimaldt.Decimal
	Reserved    decimaldt.Decimal
}
