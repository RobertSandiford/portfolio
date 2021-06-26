package routes

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo"
	"github.com/shopspring/decimal"
	"strconv"
	"gorm.io/gorm"

	"main/models"
	"main/binance"
)

func MainRoutes(db *gorm.DB, e *echo.Echo) {

	////////////////////////
	// Front Page
	////////////////////////

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "frontpage.html", nil)
	})


	////////////////////////
	// Functinoality
	////////////////////////

	e.POST("/commit", func(c echo.Context) error {

		fmt.Println("Commit")
 

		//// get the post data
		var commitPost CommitPost
		jsonDecodePostBody(c, &commitPost)

		fmt.Printf("posted assets: %+v\n", commitPost)


		//// validation stage
		validationErrors := make([]string, 0)
		var totalPercentage decimal.Decimal = decimal.NewFromFloat(0.0)
		
		for _, a := range commitPost.Assets {
			totalPercentage = totalPercentage.Add(a.Percentage)
		}

		if totalPercentage.Cmp(decimal.NewFromInt(100)) == 1 {
			validationErrors = append(validationErrors, "Desired assets percentage exceeds 100%")
		}

		if len(validationErrors) > 0 {
			r := CommitValidationErrorResponse{
				Status: "Validation Error",
				Errors: validationErrors,
			}
			return c.JSON(http.StatusOK, r)
		}

		prices := binance.GetPricedBalancesPlus()

		fmt.Printf("prices: %+v\n", prices)

		assetsMap := make(map[string]AssetPost)
		pricedBalancesMap := make(map[string]binance.BinancePricedBalance)

		for _, assetPost := range commitPost.Assets {
			assetsMap[assetPost.Symbol] = assetPost
		}

		for _, pricedBalance := range prices.PricedBalances {
			pricedBalancesMap[pricedBalance.Asset] = pricedBalance
		}

		fmt.Printf("assetsMap: %+v\n", assetsMap)
		fmt.Printf("pricedBalancesMap: %+v\n", pricedBalancesMap)

		assetsToTransact := []binance.AssetToTransact{}

		for symbol, asset := range assetsMap {
			if bal, ok := pricedBalancesMap[symbol]; ok {

				prop := asset.Percentage.Div(decimal.NewFromInt(100)).Sub(bal.Portion)
				am := prices.TotalsMap[asset.Symbol].Mul(prop)

				assetsToTransact = append(assetsToTransact, binance.AssetToTransact{
					Symbol: asset.Symbol,
					TargetPortion: asset.Percentage.Div(decimal.NewFromInt(100)),
					PortionChange: prop,
					TargetAmount: prices.TotalsMap[asset.Symbol].Mul(asset.Percentage.Div(decimal.NewFromInt(100))),
					AmountChange: am,
				})

			} else {

				prop := asset.Percentage.Div(decimal.NewFromInt(100))
				am := prices.TotalsMap[asset.Symbol].Mul(prop)

				assetsToTransact = append(assetsToTransact, binance.AssetToTransact{
					Symbol: asset.Symbol,
					TargetPortion: prop,
					PortionChange: prop,
					TargetAmount: am,
					AmountChange: am,
				})

			}
		}

		for symbol, bal := range pricedBalancesMap {
			if asset, ok := assetsMap[symbol]; !ok {

				prop := bal.Portion.Mul(decimal.NewFromInt(-1))
				am := prices.TotalsMap[asset.Symbol].Mul(prop)

				assetsToTransact = append(assetsToTransact, binance.AssetToTransact{
					Symbol: asset.Symbol,
					TargetPortion: decimal.NewFromInt(0),
					PortionChange: prop,
					TargetAmount: decimal.NewFromInt(0),
					AmountChange: am,
				})

			} 
		}

		return nil
		
		var assets []models.Asset
		for _, assetPost := range commitPost.Assets {

			percentage := assetPost.Percentage
			leverage, err := strconv.Atoi(assetPost.Leverage)
			if err != nil {
				panic("Error converting leverage to int")
			}

			asset := models.Asset{
				Symbol:     assetPost.Symbol,
				Percentage: percentage,
				Leverage:   leverage,
			}
			db.Create(&asset)

			assets = append(assets, asset)
		}

		binance.Deploy(db, assets)

		fmt.Println("Commit complete")

		//return c.Render(http.StatusOK, "sign-up-success.html", nil)
		return c.JSON(http.StatusOK, Response{Status: "OK"})
	})

	
	e.GET("/test", func(c echo.Context) error {

		r, br := binance.Buy( "BTCUSDT", decimal.NewFromFloat(10) )
		if br.Code != 0 {
			fmt.Println("Buy failed:", br.Msg)
		}


		fmt.Printf("/test: %+v\n", r)

		return c.JSON(http.StatusOK, Response{Status: "OK"})
	})

	e.GET("/balances", func(c echo.Context) error {

		pricedBalancesPlus := binance.GetPricedBalancesPlus()

		Dump(pricedBalancesPlus)

		return c.JSON(http.StatusOK, pricedBalancesPlus)
	})

	e.GET("/prices", func(c echo.Context) error {

		balances := binance.MockGetBalances()

		Dump(balances)

		return c.JSON(http.StatusOK, balances)
	})

}