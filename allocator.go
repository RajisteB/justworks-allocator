package main

import (
	"errors"
)

func Allocate(amount int, exchangeRate float64, allocationPct float64) (float64, error){
	if (amount == 0 || amount < 2) {
		amountError := errors.New("the given amount is too low to initiate a purchase, please add more funds")
		return 0, amountError
	}

	if (exchangeRate == 0) {
		exchangeRateError := errors.New("please enter a valid exchange rate")
		return 0, exchangeRateError
	}
	
	if (allocationPct == 0 || allocationPct > 1) {
		allocationPctError := errors.New("please enter a valid allocation percent")
		return 0, allocationPctError
	}
	
	allocationAmount := float64(amount) * allocationPct
	total := allocationAmount * exchangeRate
	return total, nil
}
