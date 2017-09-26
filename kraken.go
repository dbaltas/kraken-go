package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/beldur/kraken-go-api-client"
)

func formatPairTickerInfo(pair string, pairTicker krakenapi.PairTickerInfo) string {
	lastPrice, _ := strconv.ParseFloat(pairTicker.Close[0], 64)
	volume, _ := strconv.ParseFloat(pairTicker.Volume[1], 64)
	volumeAveragePrice, _ := strconv.ParseFloat(pairTicker.VolumeAveragePrice[1], 64)
	return fmt.Sprintf("%-10s %9.3f %9.3f %12.3f %9.3fM",
		pair,
		lastPrice,
		pairTicker.OpeningPrice,
		volume,
		volume*volumeAveragePrice/1000000)
}

func main() {
	api := krakenapi.New("KEY", "SECRET")

	ticker, err := api.Ticker(krakenapi.BCHEUR, krakenapi.XXBTZEUR,
		krakenapi.XXMRZEUR, krakenapi.DASHEUR)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%-10s %9s %9s %12s %10s\n",
		"Pair", "Last", "24h ago", "24h Volume", "24h EUR")
	fmt.Println(formatPairTickerInfo("BCH/EUR", ticker.BCHEUR))
	fmt.Println(formatPairTickerInfo("XBT/EUR", ticker.XXBTZEUR))
	fmt.Println(formatPairTickerInfo("XMR/EUR", ticker.XXMRZEUR))
	fmt.Println(formatPairTickerInfo("DASH/EUR", ticker.DASHEUR))
}
