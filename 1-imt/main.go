package main

import (
	"errors"
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
	_, err := fmt.Scan(&h)
	if err != nil {
		return 0, 0, err
	}

	fmt.Print("Enter your weight, kg: ")
	_, err = fmt.Scan(&w)
	if err != nil {
		return 0, 0, err
	}
	return w, h, nil
}

func calculateBMI(weight, height float64) (float64, error) {
	if weight <= 0 || height <= 0 {
		return 0, errors.New("weight or height lower or equal to zero")
	}
	hM := height / 100
	return weight / (hM * hM), nil
}

func outputResult(bmi float64) string {
	var category string
	switch {
	case bmi < 16:
		category = "High mass deficit"
	case bmi < 18.5:
		category = "Mass deficit"
	case bmi < 25:
		category = "Normal"
	case bmi < 30:
		category = "Higher than normal"
	default:
		category = "Obesity"
	}
	return category
	// fmt.Printf("Index of Body Mass: %.2f. It's %s\n", bmi, category)
}

func checkRepeatCalculation() (bool, error) {
	fmt.Println("Do you want to make another calculation? (y/n)")

	for {
		var answer string
		if _, err := fmt.Scan(&answer); err != nil {
			err = errors.New("error while reading user answer")
			return false, err
		}

		switch answer {
		case "y":
			return true, nil
		case "n":
			return false, nil
		default:
			fmt.Println("Invalid input")
		}
	}
}

func outputResult(imt float64, resultStatement string) {
	fmt.Printf("Index of Body Mass: %.2f. Its %s", imt, resultStatement)
}
