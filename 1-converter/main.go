package main

import (
	"errors"
	"fmt"
)

func main() {
	var sourceCur, targetCur string
	sourceCur = getUserCurrency(sourceCur)
	amount := getUserAmount()
	targetCur = getUserCurrency(sourceCur)
	summ, err := convert(amount, sourceCur, targetCur)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Результат конвертирования: %v", summ)
}

func convert(amount float64, from string, to string) (float64, error) {
	const USDtoEUR float64 = 0.88
	const USDtoRUB float64 = 83
	const UERtoRUB float64 = USDtoRUB / USDtoEUR

	switch {
	case from == "USD" && to == "UER":
		return amount * USDtoEUR, nil
	case from == "USD" && to == "RUB":
		return amount * USDtoRUB, nil
	case from == "UER" && to == "USD":
		return amount / USDtoEUR, nil
	case from == "UER" && to == "RUB":
		return amount * UERtoRUB, nil
	case from == "RUB" && to == "UER":
		return amount / UERtoRUB, nil
	case from == "RUB" && to == "USD":
		return amount / USDtoRUB, nil
	default:
		return 0, errors.New("Произошла исключительная ситуация")
	}
}

func getUserCurrency(currency string) string {
	var userCurrency, msg string
	for {
		switch currency {
		case "USD":
			msg = fmt.Sprint("Введите валюту (UER, RUB):")
		case "UER":
			msg = fmt.Sprint("Введите валюту (USD, RUB):")
		case "RUB":
			msg = fmt.Sprint("Введите валюту (USD, UER):")
		default:
			msg = fmt.Sprint("Введите валюту (USD, UER, RUB):")
		}

		fmt.Print(msg)
		fmt.Scan(&userCurrency)

		if userCurrency == "USD" || userCurrency == "UER" || userCurrency == "RUB" {
			if currency == userCurrency {
				fmt.Println("Ошибка: валюта не может повторятся")
				continue
			}
			return userCurrency
		}
		fmt.Println("Ошибка: валюта не определена!")
	}
}

func getUserAmount() float64 {
	var amount float64
	for {
		fmt.Print("Введите сумму: ")
		fmt.Scan(&amount)

		if amount > 0 {
			return amount
		}
		fmt.Println("Вы ввели неверное значение!")
	}
}
