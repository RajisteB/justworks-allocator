package main

import (
	"testing"
)

func TestAllocator(t *testing.T) {
	Expected := 50.0
	amount := 100
	rate := 1.0
	pct := 0.5

	result, _ := Allocate(amount, rate, pct)
	if (result != Expected) {
		t.Errorf("Result actual = %v, and Expected = %v.", result, Expected)
	}
}

func TestAllocatorZeroFunds(t *testing.T) {
	Expected := "the given amount is too low to initiate a purchase, please add more funds"
	amount := 0
	rate := 0.0025
	pct := 0.7

	_, err := Allocate(amount, rate, pct)
	if (err.Error() != Expected) {
		t.Errorf("Error actual = %v, and Expected = %v.", err.Error(), Expected)
	}
}

func TestAllocatorInvalidExchangeRate(t *testing.T) {
	Expected := "please enter a valid exchange rate"
	amount := 100
	rate := 0.0
	pct := 0.7

	_, err := Allocate(amount, rate, pct)
	if (err.Error() != Expected) {
		t.Errorf("Error actual = %v, and Expected = %v.", err.Error(), Expected)
	}
}

func TestAllocatorInvalidAllocationPct(t *testing.T) {
	Expected := "please enter a valid allocation percent"
	amount := 100
	rate := 0.0025
	pct := 1.5

	_, err := Allocate(amount, rate, pct)
	if (err.Error() != Expected) {
		t.Errorf("Error actual = %v, and Expected = %v.", err.Error(), Expected)
	}
}