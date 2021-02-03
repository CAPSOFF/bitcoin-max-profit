package main

import (
	"fmt"
	"math"
)

func calculateProfit(bitcoinPrice []float64) (profit float64) {
	maxProfit := float64(0)
	minPrice := bitcoinPrice[0]
	for _, price := range bitcoinPrice {
		maxProfit = math.Max(maxProfit, price-minPrice)
		minPrice = math.Min(minPrice, price)
	}

	return maxProfit
}

func main() {
	profit := calculateProfit([]float64{3, 2, 1, 5, 6, 2})

	fmt.Println("Profit: ", profit)
}
