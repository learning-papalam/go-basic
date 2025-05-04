package main

import "fmt"

func main() {
	var transactions []float64

	for {
		transaction := scanTranasaction()
		if transaction == 0 {
			break
		}
		transactions = append(transactions, transaction)
	}

	balance := calculateBalance(transactions)
	fmt.Printf("Ваш баланс составляет: %.2f", balance)
}

func scanTranasaction() float64 {
	var transaction float64
	fmt.Print("Введите сумму (0 для завершения):")
	fmt.Scan(&transaction)
	return transaction
}

func calculateBalance(transactions []float64) float64 {
	var balance float64
	for _, value := range transactions {
		balance += value
	}
	return balance
}
