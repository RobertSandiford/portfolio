package binance

import (
	"fmt"
	"gorm.io/gorm"
	"github.com/shopspring/decimal"

	"main/models"
)

func Deploy (db *gorm.DB, assets []models.Asset) {

	var main models.Main
	db.Select("pool").Find(&main)

	for _, asset := range assets {

		var position models.Position
		result := db.Where(&models.Position{Symbol: asset.Symbol}).First(&position)

		currentPosition, _ := decimal.NewFromString("0")
		if result.RowsAffected >= 1 {
			//
		}

		price := GetAveragePrice(asset.Symbol)
		fmt.Printf("%+v\n", price)

		amountToBuy := main.Pool.Mul(asset.Percentage.Div(decimal.NewFromFloat32(100.))).Sub(currentPosition)
		fmt.Println("amountToBuy", amountToBuy)

	}

}