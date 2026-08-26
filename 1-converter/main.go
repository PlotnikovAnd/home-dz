package main

import "fmt"

const eurFor1UsdRate float64 = 0.86
const rubFor1UsdRate float64 = 84.37
const eurToRubRate float64 = rubFor1UsdRate / eurFor1UsdRate

func main() {
	fmt.Println("rub for 1 eur:", eurToRubRate)
}
