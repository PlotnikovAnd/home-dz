package main

import (
	"fmt"
)

func main() {

	fmt.Println("__ Body Mass Calculator __")

	userWeight, userHeight := getUserInput()
	resultIMT := calculateIMT(userWeight, userHeight)

	var resultStatement string
	switch {
	case resultIMT < 16:
		resultStatement = "High mass deficit"
	case resultIMT < 18.5:
		resultStatement = "Mass deficit"
	case resultIMT < 25:
		resultStatement = "Normal"
	case resultIMT < 30:
		resultStatement = "Higher than normal"
	default:
		resultStatement = "Obesity degree"
	}

	outputResult(resultIMT, resultStatement)
}

func getUserInput() (userWeight, userHeight float64) {

	fmt.Print("Enter your height, cm: ")
	_, err := fmt.Scan(&userHeight)
	if err != nil {
		fmt.Println("invalid input:", err)
		return 0, 0
	}

	fmt.Print("Enter your weight, kg: ")
	_, err = fmt.Scan(&userWeight)
	if err != nil {
		fmt.Println("invalid input:", err)
		return 0, 0
	}
	return userWeight, userHeight
}

func calculateIMT(weight, height float64) float64 {
	return weight / (height / 100 * height / 100)
}

func outputResult(imt float64, resultStatement string) {
	fmt.Printf("Index of Body Mass: %.2f. Its %s", imt, resultStatement)
}
