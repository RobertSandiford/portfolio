package binance

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"github.com/shopspring/decimal"

	//neturl "net/url"
	//"encoding/hex"
	//"hash"
	//"crypto/sha256"
	//"crypto/hmac"
)

var binanceApiKey string
var binanceApiSecret string

func Init() {
	binanceApiKey = os.Getenv("binance_api_key")
	binanceApiSecret = os.Getenv("binance_api_secret")
}

func get(url, queryString, task string) (responseBody string) {
	fmt.Println("GET this address", url+queryString)
	res, err := http.Get(url + queryString)
	if err != nil {
		panic("Error doing GET while " + task)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic("Error reading response body while " + task)
	}
	fmt.Println(string(body))

	return string(body)
}

func GetExchangeInfo() {
	task := "GETing exchange info"
	url := "https://api.binance.com/api/v3/exchangeInfo"

	get(url, "", task)
}

type Fourth struct {
	Four_one int    `json:"four_one"`
	Four_two string `json:"four_two"`
}
type MockResponse struct {
	First  int    `json:"first"`
	Second string `json:"second"`
	Third  string `json:"third"`
	Fourth Fourth `json:"fourth"`
}

func MockRequest() {
	jSource := `{"first":1,"second":"abc","third":"foo","fourth":{"four_one":41,"four_two":"bar"}}`
	j := []uint8(jSource)

	fmt.Println(string(j))

	var data MockResponse

	json.Unmarshal(j, &data)

	fmt.Printf("%+v\n", data)
}

func GetAveragePrice(pair string) (price decimal.Decimal) {
	task := "GETing averge price"

	url := "https://api.binance.com/api/v3/avgPrice"
	params := []Param{
		Param{Name: "symbol", Value: pair},
	}

	queryString := makeQueryString(params)

	//req, _ := http.NewRequest("GET", ur+queryString, ni)
	//res, _ := http.DefaultClient.Do(re)
	//fmt.Println("Get this address", url+queryString)

	res, err := http.Get(url + queryString)
	if err != nil {
		panic("Error " + task)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic("Error reading body after " + task)
	}
	//fmt.Printf("%T\n", body)
	//fmt.Println(string(body))

	var data GetAveragePriceData
	JsonDecodeBytes(body, &data)

	//json.Unmarshal([]byte(body), &data)

	//fmt.Printf("%+v\n", data)

	return data.Price
}

func GetBalances() (balances []BinanceBalance) {
	task := "GETing binance balances"

	url := "https://api.binance.com/api/v3/account"

	params := []Param{
		Param{ Name: "timestamp", Value: MiliTimestampString() },
	}

	//queryString := makeQueryString(params)

	params = signParams(params, "GET")

	body := binanceGet(url, params, task)

	var data BinanceAccount
	JsonDecodeBytes(body, &data)

	fmt.Printf("%+v\n", data)

	return data.Balances
}

func MockGetBalances() (balances []BinanceBalance) {
	return []BinanceBalance{
		BinanceBalance{ Asset: "BTC", Free: decimal.NewFromFloat(0.002), Locked: decimal.NewFromFloat(0.0) },
		BinanceBalance{ Asset: "DOT", Free: decimal.NewFromFloat(30.45), Locked: decimal.NewFromFloat(0.0) },
		BinanceBalance{ Asset: "COTI", Free: decimal.NewFromFloat(15.4), Locked: decimal.NewFromFloat(0.0) },
	}
}

func GetPricedBalancesPlus() (prices BinancePricedBalancesPlus) {

	balances := MockGetBalances()
	//balances := MockBalances()

	var totalBtc decimal.Decimal

	for _, balance := range balances {
		if balance.Asset == "BTC" {

			//fmt.Println("Asset is BTC")

			pricedBalance := BinancePricedBalance{
				Asset: balance.Asset,
				Free: balance.Free,
				Locked: balance.Locked,
				Total: balance.Free.Add(balance.Locked),
			}
			pricedBalance.Price = decimal.NewFromInt(1)
			pricedBalance.Value = pricedBalance.Total

			// save
			prices.PricedBalances = append(prices.PricedBalances, pricedBalance)
			totalBtc = totalBtc.Add(pricedBalance.Value)

		} else {

			//fmt.Println("Asset is NOT BTC")

			price := GetAveragePrice(balance.Asset + "BTC")

			pricedBalance := BinancePricedBalance{
				Asset: balance.Asset,
				Free: balance.Free,
				Locked: balance.Locked,
				Total: balance.Free.Add(balance.Locked),
				Price: price,
			}
			pricedBalance.Value = pricedBalance.Total.Mul(price)

			// save
			prices.PricedBalances = append(prices.PricedBalances, pricedBalance)
			totalBtc = totalBtc.Add(pricedBalance.Value)

		}
		
	}

	for k, pb := range prices.PricedBalances {
		pb.Portion = pb.Value.Div(totalBtc)
		prices.PricedBalances[k] = pb
	}

	btcGbpPrice := GetAveragePrice("BTCGBP")
	btcUsdtPrice := GetAveragePrice("BTCUSDT")

	prices.Prices = []Price{
		Price{ Pair: "BTCGBP", Price: btcGbpPrice },
		Price{ Pair: "BTCUSDT", Price: btcUsdtPrice },
	}
	prices.Totals = []Total{
		Total{ Asset: "BTC", Value: totalBtc },
		Total{ Asset: "GBP", Value: totalBtc.Mul(btcGbpPrice) },
		Total{ Asset: "USDT", Value: totalBtc.Mul(btcUsdtPrice) },
	}
	prices.TotalsMap = map[string]decimal.Decimal{
		"BTC": totalBtc,
		"GBP": totalBtc.Mul(btcGbpPrice),
		"USDT": totalBtc.Mul(btcUsdtPrice),
	}

	/*prices = BinancePricedBalancesPlus{
		Total: total,
		PricedBalances: pricedBalances,
	}*/

	return prices
}


func Buy(pair string, amount decimal.Decimal) (orderResponse OrderResponseFull, errorResponse BinanceResponse) {
	task := "POSTing standard order"

	url := "https://api.binance.com/api/v3/order"

	orderResponse, errorResponse = BuyInternal(url, pair, amount, task)

	return orderResponse, errorResponse
}
func BuyTest(pair string, amount decimal.Decimal) (orderResponse OrderResponseFull, errorResponse BinanceResponse) {
	task := "POSTing test standard order"

	url := "https://api.binance.com/api/v3/order/test"

	orderResponse, errorResponse = BuyInternal(url, pair, amount, task)

	return orderResponse, errorResponse
}
func BuyInternal(url, pair string, amount decimal.Decimal, task string) (orderResponse OrderResponseFull, errorResponse BinanceResponse) {

	//fmt.Println("url", url)

	params := []Param{
		Param{ Name: "symbol", Value: pair },
		Param{ Name: "side", Value: "BUY" },
		Param{ Name: "type", Value: "MARKET" },
		//Param{ Name: "quantity", Value: amount.String() },
		Param{ Name: "quoteOrderQty", Value: amount.String() },
		Param{ Name: "timestamp", Value: MiliTimestampString() },
	}

	params = signParams(params, "POST")

	body := binancePost(url, params, task)

	// check for error responses
	var br BinanceResponse
	JsonDecodeBytes(body, &br)
	//fmt.Printf("%+v\n", br)

	if br.Code != 0 {
		return OrderResponseFull{}, br
	}

	// else
	var or OrderResponseFull
	JsonDecodeBytes(body, &or)

	fmt.Printf("BuyInternal: %+v\n", or)

	return or, BinanceResponse{}
}


func GetPricedBalances() (pricedBalances []BinancePricedBalance) {
	pricedBalancesPlus := GetPricedBalancesPlus()
	return pricedBalancesPlus.PricedBalances
}
