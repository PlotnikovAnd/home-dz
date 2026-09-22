package main


import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func scanTransactions() ([]float64, error) {
	transactions := []float64{}
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		if line == "exit" {
			break
		}

		trans, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Printf("invalid input: %v\n", err)
			continue
		}
		transactions = append(transactions, trans)

	}

	err := scanner.Err()
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func calculateBalance(transactions []float64) float64 {
	balance := 0.0
	for _, t := range transactions {
		balance += t
	}
	return balance
}

func main() {
	tr1 := make([]int, 0, 3)
	tr1 = append(tr1, 1)
	tr1 = append(tr1, 2)
	// tr1 = append(tr1, 3)
	fmt.Printf("%v, size=%d, cap=%d\n", tr1, len(tr1), cap(tr1))
	// tr1 := []int{1, 2, 3}
	// tr2 := []int{4, 5, 6}
	// tr1 = append(tr1, tr2...) // ... - is unpack
	// fmt.Println(tr1)
	arr := [3]int{1, 2, 3}
	b := arr
	b[0] = 100
	fmt.Println(arr, b)

	transactions, err := scanTransactions()
	if err != nil {
		fmt.Printf("Error while scanning transactions: %v\n", err)
	}
	balance := calculateBalance(transactions)
	fmt.Printf("Balance = %v\n", balance)
}
