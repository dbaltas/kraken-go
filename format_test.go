package main

import (
	"github.com/beldur/kraken-go-api-client"
	"testing"
)

func TestFormat(t *testing.T) {
	pairTicker := new(krakenapi.PairTickerInfo)
	pair := "BCH/EUR"
	var close = []string{"380", "384"}
	var volume = []string{"101000", "102000"}
	var volumeAveragePrice = []string{"379", "102383"}
	pairTicker.Close = close
	pairTicker.Volume = volume
	pairTicker.VolumeAveragePrice = volumeAveragePrice
	pairTicker.OpeningPrice = 382.34
	output := formatPairTickerInfo(pair, *pairTicker)
	expected := "BCH/EUR      380.000   382.340   102000.000 10443.066M"
	if output != expected {
		t.Errorf(`%s\n%s\n`, output, expected)
	}
}
