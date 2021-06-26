package binance

import (
	"github.com/shopspring/decimal"
)

type AssetToTransact struct {
	Symbol string
	TargetPortion decimal.Decimal
	PortionChange decimal.Decimal
	TargetAmount decimal.Decimal
	AmountChange decimal.Decimal
}