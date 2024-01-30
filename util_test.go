package main

import (
	"reflect"
	"testing"
)

func TestHasArbitrageOpportunity(t *testing.T) {
	type Input struct {
		currencies map[string]map[string]float64
		target     string
	}
	type Case struct {
		title  string
		input  Input
		output bool
	}

	cases := []Case{
		{
			title: "No Arbitrage Opportunity",
			input: Input{
				currencies: map[string]map[string]float64{
					"USD": {"EUR": 0.9},
					"EUR": {"USD": 1.1},
				},
				target: "USD",
			},
			output: false,
		},
		{
			title: "Direct Arbitrage Opportunity",
			input: Input{
				currencies: map[string]map[string]float64{
					"USD": {"EUR": 0.9},
					"EUR": {"USD": 1.2},
				},
				target: "USD",
			},
			output: true,
		},
		{
			title: "Indirect Arbitrage Opportunity",
			input: Input{
				currencies: map[string]map[string]float64{
					"USD": {"EUR": 0.9},
					"EUR": {"JPY": 160},
					"JPY": {"USD": 0.008},
				},
				target: "USD",
			},
			output: true,
		},
		{
			title: "No Arbitrage Opportunity, Indirectly",
			input: Input{
				currencies: map[string]map[string]float64{
					"USD": {"EUR": 0.85},
					"EUR": {"JPY": 130},
					"JPY": {"USD": 0.0075},
				},
				target: "USD",
			},
			output: false,
		},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			result := hasArbitrageOpportunity(c.input.currencies, c.input.target)
			if result != c.output {
				t.Errorf("%s: expected %v, got %v", c.title, c.output, result)
			}
		})
	}
}

func TestTraverseCurrencies(t *testing.T) {
	type Input struct {
		currencies map[string]map[string]float64
		traversed  map[string]bool
		current    string
		target     string
		value      float64
	}
	type Case struct {
		title  string
		input  Input
		output bool
	}

	cases := []Case{
		{
			title: "No Arbitrage Opportunity",
			input: Input{
				currencies: map[string]map[string]float64{
					"USD": {"EUR": 0.9},
					"EUR": {"USD": 1.1},
				},
				traversed: make(map[string]bool),
				current:   "USD",
				target:    "USD",
				value:     1,
			},
			output: false,
		},
		{
			title: "Direct Arbitrage Opportunity",
			input: Input{
				currencies: map[string]map[string]float64{
					"USD": {"EUR": 0.9},
					"EUR": {"USD": 1.2},
				},
				traversed: make(map[string]bool),
				current:   "USD",
				target:    "USD",
				value:     1,
			},
			output: true,
		},
		{
			title: "Indirect Arbitrage Opportunity",
			input: Input{
				currencies: map[string]map[string]float64{
					"USD": {"EUR": 0.9},
					"EUR": {"JPY": 160},
					"JPY": {"USD": 0.008},
				},
				traversed: make(map[string]bool),
				current:   "USD",
				target:    "USD",
				value:     1,
			},
			output: true,
		},
		{
			title: "No Arbitrage Opportunity, Indirectly",
			input: Input{
				currencies: map[string]map[string]float64{
					"USD": {"EUR": 0.85},
					"EUR": {"JPY": 130},
					"JPY": {"USD": 0.0075},
				},
				traversed: make(map[string]bool),
				current:   "USD",
				target:    "USD",
				value:     1,
			},
			output: false,
		},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			result := traverseCurrencies(c.input.currencies, c.input.traversed, c.input.current, c.input.target, c.input.value)
			if result != c.output {
				t.Errorf("%s: expected %v, got %v", c.title, c.output, result)
			}
		})
	}
}

func TestCloneTraversed(t *testing.T) {
	type Case struct {
		title  string
		input  map[string]bool
		output map[string]bool
	}

	cases := []Case{
		{
			title:  "Empty",
			input:  map[string]bool{},
			output: map[string]bool{},
		},
		{
			title:  "USD",
			input:  map[string]bool{"USD": true},
			output: map[string]bool{"USD": true},
		},
		{
			title:  "USD, EUR",
			input:  map[string]bool{"USD": true, "EUR": false},
			output: map[string]bool{"USD": true, "EUR": false},
		},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			output := cloneTraversed(c.input)

			if !reflect.DeepEqual(output, c.output) {
				t.Errorf("Expected %v, got %v", c.output, output)
			}
		})
	}
}

func TestCloneTraversedImmutability(t *testing.T) {
	type Case struct {
		title    string
		input    map[string]bool
		modifier func(map[string]bool)
	}

	cases := []Case{
		{
			title: "Empty",
			input: map[string]bool{},
			modifier: func(original map[string]bool) {
				original["NewKey"] = true
			},
		},
		{
			title: "Single Element",
			input: map[string]bool{"USD": true},
			modifier: func(original map[string]bool) {
				original["USD"] = false
			},
		},
		{
			title: "Multiple Elements",
			input: map[string]bool{"USD": true, "EUR": false},
			modifier: func(original map[string]bool) {
				original["JPY"] = true
			},
		},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			original := make(map[string]bool)
			for key, value := range c.input {
				original[key] = value
			}

			cloned := cloneTraversed(original)
			c.modifier(original)

			if reflect.DeepEqual(original, cloned) {
				t.Errorf("Expected difference, but received two %v", original)
			}
		})
	}
}
