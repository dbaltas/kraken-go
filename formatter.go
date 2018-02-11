package main

import (
	"fmt"
	"strconv"

	"github.com/beldur/kraken-go-api-client"
)

func formatPairTickerInfo(pair string, pairTicker krakenapi.PairTickerInfo) string {
	lastPrice, _ := strconv.ParseFloat(pairTicker.Close[0], 64)
	volume, _ := strconv.ParseFloat(pairTicker.Volume[1], 64)
	low, _ := strconv.ParseFloat(pairTicker.Low[1], 64)
	high, _ := strconv.ParseFloat(pairTicker.High[1], 64)
	volumeAveragePrice, _ := strconv.ParseFloat(pairTicker.VolumeAveragePrice[1], 64)
	return fmt.Sprintf("%-10s %9.3f %9.3f %9.3f %12.3f %9.3fM",
		pair,
		lastPrice,
		low,
		high,
		volume,
		volume*volumeAveragePrice/1000000)
}
