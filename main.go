package main

import (
	"errors"
	"fmt"
	"strconv"
)


var exchangeRatesURL = "https://api.coinbase.com/v2/exchange-rates?currency=USD"

func main() {
	// Total amount to allocate
	amount := 100

	// 70% allocation to bitcoin
	btcAllocationPct := .7

	// 30% allocation to ethereum
	ethAllocationPct := .3

	// Get exchange rates
	currencyData := GetRates(exchangeRatesURL)
	btcRate, btcErr := strconv.ParseFloat(currencyData.Data.Rates["BTC"], 32)
	ethRate, ethErr := strconv.ParseFloat(currencyData.Data.Rates["ETH"], 32)

	if (btcErr != nil || ethErr != nil) {
		HandleProgramError("Error: ", errors.New("cannot parse exchange rate into type float"))
	}

	// Allocate for Bitcoin
	btcResult, err := Allocate(amount, btcRate, btcAllocationPct)
	if (err != nil) {
		HandleProgramError(err.Error(), err)
	}
	fmt.Printf("You can purchase %v amount of Bitcoin with %s dollars\n", btcResult, fmt.Sprint(float64(amount) * btcAllocationPct))
	
	// Allocate for Ethereum
	ethResult, err := Allocate(amount, ethRate, ethAllocationPct)
	if (err != nil) {
		HandleProgramError(err.Error(), err)
	}
	fmt.Printf("You can purchase %v amount of Ethereum with %s dollars\n", ethResult, fmt.Sprint(float64(amount) * ethAllocationPct))
}
