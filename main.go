package main

import (
	"fmt"
	"math"
)

// package is used for decompose and incapsulate
// for each file it should have package on top
// packages could be:
// - our packages
// - main package
// - standard packages
// - external packages

// go.mod?
// 1 module = 1 app
func main() {
	var userHeight = 1.8
	var userKg float64 = 100
	var IMT = userKg / math.Pow(userHeight, 2)
	fmt.Println(IMT)
}
