# Currency Arbitrage Detection

## Overview
A simple implementation of functionality to detect arbitrage opportunities in currency exchange rates. Arbitrage exists if there is a cycle in the currency exchange rates that starts and ends with the same currency and results in a profit.

## Functions

### `hasArbitrageOpportunity`

```go
func hasArbitrageOpportunity(currencies map[string]map[string]float64, target string) bool
```

- **Description**: Determines whether there is an arbitrage opportunity starting from a given target currency.
- **Parameters**:
    - `currencies`: A map of currency pairs and their exchange rates.
    - `target`: The currency from which to start looking for an arbitrage opportunity.
- **Returns**: `true` if there is an arbitrage opportunity; otherwise, `false`.

### `traverseCurrencies`

```go
func traverseCurrencies(currencies map[string]map[string]float64, traversed map[string]bool, current, target string, value float64) bool
```

- **Description**: Recursively traverses through currency pairs to find arbitrage opportunities.
- **Parameters**:
    - `currencies`: A map of currency pairs and their exchange rates.
    - `traversed`: A map to keep track of visited currencies during traversal.
    - `current`: The current currency being evaluated.
    - `target`: The target currency for finding an arbitrage cycle.
    - `value`: The cumulative product of exchange rates along the current path.
- **Returns**: `true` if an arbitrage opportunity is found; otherwise, `false`.

### `cloneTraversed`

```go
func cloneTraversed(traversed map[string]bool) map[string]bool
```

- **Description**: Creates a clone of the map used to track visited currencies.
- **Parameters**:
    - `traversed`: A map indicating currencies that have been visited.
- **Returns**: A new map that is a clone of the `traversed` map.

## Usage

Example:

```go
currencies := map[string]map[string]float64{
    "USD": {"EUR": 0.9},
    "EUR": {"JPY": 160},
    "JPY": {"USD": 0.008},
}

arbitrage := hasArbitrageOpportunity(currencies, "USD")
if arbitrage {
    fmt.Println("Arbitrage opportunity detected!")
} else {
    fmt.Println("No arbitrage opportunity.")
}
```

To run the program, execute the following command:

```bash
go run .
```

## Testing

To run the tests, execute the following command:

```bash
go test
```