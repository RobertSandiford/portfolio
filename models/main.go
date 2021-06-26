package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Main struct {
	gorm.Model
	Pool     decimal.Decimal `sql:"type:decimal(18,18);"`
	MainCoin string
}