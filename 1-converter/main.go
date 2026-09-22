package main

import "fmt"


func main() {
	m := map[string]map[string]float64{
		"eur": {
			"usd": 1.0 / 0.877,
			"rub": 101.0,
		},
		"usd": {
			"eur": 0.86,
			"rub": 87.7,
		},
		"rub": {
			"eur": 1.0 / 101.0,
			"usd": 1.0 / 87.7,
		},
	}

	var from, to string
	var quantity float64

	fmt.Print("Enter currency ticker FROM which to convert: ")
	fmt.Scan(&from)
	fmt.Print("Enter currency ticker TO which to convert: ")
	fmt.Scan(&to)
	fmt.Print("Enter quantity: ")
	fmt.Scan(&quantity)

	result := m[from][to] * quantity
	fmt.Println("You will get:", result)
}
