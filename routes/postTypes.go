package routes

import (
	"github.com/shopspring/decimal"
)

/*
type AssetPost struct {
	Symbol     string
	Percentage string
	Leverage   string
}
*/

type AssetPost struct {
	Symbol     string `json:"symbol"`
	Percentage decimal.Decimal `json:"percentage"`
	Leverage   string `json:"leverage"`
}

type CommitPost struct {
	Assets []AssetPost `json:"assets"`
}
