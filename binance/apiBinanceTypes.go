package binance

import (
	"github.com/shopspring/decimal"
)

type Param struct {
	Name  string
	Value string
}

type BinanceResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type GetAveragePriceData struct {
	Mins  int
	Price decimal.Decimal
}

type BinanceBalance struct {
	Asset string  `json:"asset"`
	Free decimal.Decimal `json:"free"`
	Locked decimal.Decimal `json:"locked"`
}

type BinanceAccount struct {
	CanTrade  bool `json:"canTrade"`
	CanDeposit  bool `json:"canDeposit"`
	UpdateTime int `json:"updateTime"`
	AccountType string `json:"accountType"`
	Balances []BinanceBalance `json:"balances"`
	Permissions []string `json:"permissions"`
}

type Total struct {
	Asset string `json:"asset"`
	Value decimal.Decimal `json:"value"`
}

type Price struct {
	Pair string `json:"pair"`
	Price decimal.Decimal `json:"price"`
}

type BinancePricedBalance struct {
	Asset string  `json:"asset"`
	Free decimal.Decimal `json:"free"`
	Locked decimal.Decimal `json:"locked"`
	Total decimal.Decimal `json:"total"`
	Price decimal.Decimal `json:"price"`
	Value decimal.Decimal `json:"value"`
	Portion decimal.Decimal `json:"portion"`
}

type BinancePricedBalancesPlus struct {
	Totals []Total `json:"totals"`
	TotalsMap map[string]decimal.Decimal `json:"totalsMap"`
	Prices []Price `json:"prices"`
	PricedBalances []BinancePricedBalance `json:"pricedBalances"`
}

type Fill struct {
	Price           string `json:"price"`
	Qty             string `json:"qty"`
	Commission      string `json:"commission"`
	CommissionAsset string `json:"commissionAsset"`
}

type OrderResponseFull struct {
	Symbol              string `json:"symbol"`
	OrderId             int    `json:"orderId"`
	OrderListId         int    `json:"orderListId"`
	ClientListId        string `json:"clientListId"`
	TransactTime        string `json:"transactTime"`
	Price               string `json:"price"`
	OrigQty             string `json:"origQty"`
	ExecutedQty         string `json:"executedQty"`
	CummulativeQuoteQty string `json:"cummulativeQuoteQty"`
	Status              string `json:"status"`
	TimeInForce         string `json:"timeInForce"`
	OrderType           string `json:"type"`
	Side                string `json:"side"`
	Fills               []Fill `json:"fills"`
}

type RateLimit struct {
}
type ExchangeFilters struct {
}
type Filter struct {
}
type Symbol struct {
	Symbol                     string   `json:"symbol"`
	Status                     string   `json:"status"`
	BaseAsset                  string   `json:"baseAsset"`
	BaseAssetPrecision         int      `json:"baseAssetPrecision"`
	QuoteAsset                 string   `json:"quoteAsset"`
	QuotePrecision             int      `json:"quotePrecision"`
	QuoteAssetPrecision        int      `json:"quoteAssetPrecision"`
	BaseCommissionPrecision    int      `json:"baseCommissionPrecision"`
	QuoteCommissionPrecision   int      `json:"quoteCommissionPrecision"`
	OrderTypes                 []string `json:"orderTypes"`
	IcebergAllowed             bool     `json:"icebergAllowed"`
	OcoAllowed                 bool     `json:"ocoAllowed"`
	QuoteOrderQtyMarketAllowed bool     `json:"quoteOrderQtyMarketAllowed"`
	IsSpotTradingAllowed       bool     `json:"isSpotTradingAllowed"`
	IsMarginTradingAllowed     bool     `json:"isMarginTradingAllowed"`
	Filters                    []Filter `json:"filters"`
	Permissions                []string `json:"permissions"`
}

type DepthResponse struct {
	Timezone        string            `json:"timezone"`
	ServerTime      int               `json:"serverTime"`
	RateLimits      []RateLimit       `json:"rateLimits"`
	ExchangeFilters []ExchangeFilters `json:"exchangeFilters"`
	symbols         []Symbol          `json:"symbols"`
}
