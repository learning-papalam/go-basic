package main

import (
	"fmt"
	"strings"
)

type currencyExchange map[string]map[string]float32
type userData map[string]string

func main() {
	var base currencyExchange
	var amount float32
	userCurrency := userData{}

	fabrica(&base)

	getUserCurrence(&userCurrency, &base)
	amount = getUserAmount()
	getUserCurrence(&userCurrency, &base)
	calculatedAmound := calculate(amount, &userCurrency, &base)

	fmt.Printf("В результате Вы получите: %0.2f", calculatedAmound)
}

func fabrica(c *currencyExchange) {
	const USDtoRUB float32 = 88
	const USDtoUER float32 = 0.90
	const UERtoRUB float32 = USDtoRUB / USDtoUER

	*c = currencyExchange{
		"USD": {
			"RUB": USDtoRUB,
			"EUR": USDtoUER,
		},
		"EUR": {
			"RUB": UERtoRUB,
			"USD": 1 / USDtoUER,
		},
		"RUB": {
			"USD": 1 / USDtoRUB,
			"EUR": 1 / UERtoRUB,
		},
	}
}

func getUserCurrence(u *userData, b *currencyExchange) {
	var inputCurrency string
	avalibleCurrency := []string{}
	isCorrecCurrency := false
	if len(*u) == 0 {
		for currency := range *b {
			avalibleCurrency = append(avalibleCurrency, currency)
		}
	} else {
		from := (*u)["from"]
		for currency := range *b {
			if from == currency {
				continue
			}
			avalibleCurrency = append(avalibleCurrency, currency)
		}
	}
Loop:
	for {
		fmt.Printf("Выберите валюту %v: ", strings.Join(avalibleCurrency, ","))
		fmt.Scan(&inputCurrency)

		inputCurrency = strings.ToUpper(inputCurrency)
		for currency := range *b {
			if currency == inputCurrency {
				isCorrecCurrency = true
			}
		}
		if isCorrecCurrency {
			if len(*u) == 0 {
				(*u)["from"] = inputCurrency
				break Loop
			} else {
				from := (*u)["from"]
				if from == inputCurrency {
					fmt.Println("Валюта не может совпадать!")
					continue Loop
				}
				(*u)["to"] = inputCurrency
				break Loop
			}
		}
		fmt.Println("Вы ошиблись при вводе валюты!")

	}

}

func getUserAmount() float32 {
	var userAmount float32
	for {
		fmt.Print("Введите количество: ")
		fmt.Scan(&userAmount)
		if userAmount > 0 {
			return userAmount
		}
		fmt.Println("Ошибка при вводе, попробуйте еще раз...")
	}
}

func calculate(amount float32, u *userData, b *currencyExchange) float32 {
	var from string = (*u)["from"]
	var to string = (*u)["to"]
	return amount * (*b)[from][to]
}
