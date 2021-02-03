package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
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

func readFileString(filepath string) string {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

func main() {
	var totalProfit float64
	fileContainString := readFileString("gistfile1.txt")
	splitedInput := strings.Split(fileContainString, " ")

	for _, parsedInput := range splitedInput {
		var tempProfit []float64
		for _, val := range parsedInput {
			floatVal, err := strconv.ParseFloat(string(val), 64)
			if err != nil {
				log.Fatal(err)
			}

			tempProfit = append(tempProfit, floatVal)
		}

		totalProfit += calculateProfit(tempProfit)
	}

	fmt.Println(totalProfit)

}
