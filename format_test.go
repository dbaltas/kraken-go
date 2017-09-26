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
	var high = []string{"385", "387"}
	var low = []string{"372", "370"}
	pairTicker.Close = close
	pairTicker.Volume = volume
	pairTicker.VolumeAveragePrice = volumeAveragePrice
	pairTicker.High = high
	pairTicker.Low = low

	output := formatPairTickerInfo(pair, *pairTicker)
	expected := "BCH/EUR      380.000   370.000   387.000   102000.000 10443.066M"
	if output != expected {
		t.Errorf(`%s\n%s\n`, output, expected)
	}
}
