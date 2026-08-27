package main

import "fmt"

const eurFor1UsdRate float64 = 0.86
const rubFor1UsdRate float64 = 84.37
const eurToRubRate float64 = rubFor1UsdRate / eurFor1UsdRate

func parseUserInput() float64 {
	var input float64
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("wrong input, error:", err)
	}
	return input
}

func calculateCurrency(amount float64, from string, to string) float64 {
	return 1.0
}

func main() {
	result := calculateCurrency(100, "usd", "rub")
	fmt.Println("", result)
}
