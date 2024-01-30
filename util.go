package main

func hasArbitrageOpportunity(currencies map[string]map[string]float64, target string) bool {
	return traverseCurrencies(currencies, make(map[string]bool), target, target, 1)
}

func traverseCurrencies(currencies map[string]map[string]float64, traversed map[string]bool, current, target string, value float64) bool {
	if traversed[current] {
		if current == target && value > 1 {
			return true
		}
		return false
	}

	traversed[current] = true
	for next, rate := range currencies[current] {
		if traverseCurrencies(currencies, cloneTraversed(traversed), next, target, value*rate) {
			return true
		}
	}

	return false
}

func cloneTraversed(traversed map[string]bool) map[string]bool {
	clone := make(map[string]bool)
	for key, value := range traversed {
		clone[key] = value
	}
	return clone
}
