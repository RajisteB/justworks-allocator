package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type CoinbaseResponse struct {
	Data Currencies
}

type Currencies struct {	
	Currency string
	Rates map[string]string
}


func GetRates(url string) CoinbaseResponse {
	var currencyData CoinbaseResponse
	
	resp, err := http.Get(url)
	if err != nil {
		HandleProgramError("Error making GET request: %s\n", err)
	}

	defer resp.Body.Close()


	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		HandleProgramError("Client cannot read response body %s\n", err)
	}

	err = json.Unmarshal(resBody, &currencyData)
	if err != nil {
		HandleProgramError("Client cannot parse body %s\n", err)
	}

	return currencyData
}
