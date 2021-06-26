package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Asset struct {
	gorm.Model
	Symbol     string
	Percentage decimal.Decimal `sql:"type:decimal(3,2);"`
	Leverage   int
}
