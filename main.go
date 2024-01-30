package main

import "fmt"

// Write a function that takes as input a set of currency
// pair exchange rates and determines whether there is an
// arbitrage opportunity or not.

// Example:
// 'USD', 'EUR', .9
// 'EUR', 'JPY', 160
// 'JPY', 'USD', .008
//
// 1 USD = .9 EUR = 144 JPY = 1.152 USD

func main() {
	currencies := map[string]map[string]float64{
		"USD": {"EUR": 0.9},
		"EUR": {"JPY": 160},
		"JPY": {"USD": 0.008},
	}

	target := "USD"

	if hasArbitrageOpportunity(currencies, target) {
		fmt.Printf("There is an arbitrage opportunity for %v\n", target)
	} else {
		fmt.Printf("No arbitrage opportunity for %v arbitrage\n", target)
	}
}
