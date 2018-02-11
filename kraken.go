package main

import (
	"fmt"
	"log"

	"github.com/beldur/kraken-go-api-client"
)

func main() {
	api := krakenapi.New("KEY", "SECRET")

	pairs := []string{
		krakenapi.XETHZEUR,
		krakenapi.BCHEUR,
		krakenapi.XXBTZEUR,
		krakenapi.XXMRZEUR,
		krakenapi.DASHEUR,
	}

	ticker, err := api.Ticker(pairs...)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%-10s %9s %9s %9s %12s %10s\n",
		"Pair", "Last", "24h low", "24h high", "24h Volume", "24h EUR")

	fmt.Println(formatPairTickerInfo("ETH/EUR", ticker.XETHZEUR))
	fmt.Println(formatPairTickerInfo("BCH/EUR", ticker.BCHEUR))
	fmt.Println(formatPairTickerInfo("XBT/EUR", ticker.XXBTZEUR))
	fmt.Println(formatPairTickerInfo("XMR/EUR", ticker.XXMRZEUR))
	fmt.Println(formatPairTickerInfo("DASH/EUR", ticker.DASHEUR))
}
