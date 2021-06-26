package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Trade struct {
	gorm.Model
	Symbol   string
	Against  string
	Action   string
	Amount   decimal.Decimal `sql:"type:decimal(18,18);"`
	Leverage int
}