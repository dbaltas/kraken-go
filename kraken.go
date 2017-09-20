package main

import (
	"fmt"
	"log"

	"github.com/beldur/kraken-go-api-client"
)

func main() {
	api := krakenapi.New("KEY", "SECRET")

	ticker, err := api.Ticker(krakenapi.BCHEUR, krakenapi.XXBTZEUR,
		krakenapi.XXMRZEUR, krakenapi.DASHEUR)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("BCH/EUR")
	fmt.Println(ticker.BCHEUR.OpeningPrice)
	fmt.Println(ticker.BCHEUR.Close)
	fmt.Println(ticker.BCHEUR.Volume)

	fmt.Println()
	fmt.Println("XBT/EUR")
	fmt.Println(ticker.XXBTZEUR.OpeningPrice)
	fmt.Println(ticker.XXBTZEUR.Close)
	fmt.Println(ticker.XXBTZEUR.Volume)

	fmt.Println()
	fmt.Println("XMR/EUR")
	fmt.Println(ticker.XXMRZEUR.OpeningPrice)
	fmt.Println(ticker.XXMRZEUR.Close)
	fmt.Println(ticker.XXMRZEUR.Volume)

	fmt.Println()
	fmt.Println("DASH/EUR")
	fmt.Println(ticker.DASHEUR.OpeningPrice)
	fmt.Println(ticker.DASHEUR.Close)
	fmt.Println(ticker.DASHEUR.Volume)
}
