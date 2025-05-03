package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Произошла ошибка:\n", r)
		}
	}()

	fmt.Println("__ Калькулятор индекса массы тела __")

	for true {
		userHeigth, userWidth := getUserInput()
		IMT, err := calculateIMT(userHeigth, userWidth)
		if err != nil {
			panic(fmt.Sprintf("Ошибка: %s", err.Error()))
			// fmt.Println(err)
			// continue
		}
		outputResult(IMT)

		againCalculate := userResponseCalculate()
		if !againCalculate {
			break
		}

	}

}

func outputResult(IMT float64) {
	result := fmt.Sprintf("Индекс массы тела: %.2f", IMT)
	fmt.Println(result)
	switch {
	case IMT < 16:
		fmt.Println("У вас сильный дефицит массы тела")
		break
	case IMT < 18.5:
		fmt.Println("У вас дефицит массы тела")
	case IMT < 25:
		fmt.Println("У вас вес в норме")
	case IMT < 30:
		fmt.Println("У вас предожирение")
	default:
		fmt.Println("У вас ожирение")
	}
}

func calculateIMT(userHeigth float64, userWidth float64) (float64, error) {
	const Power = 2

	if userHeigth <= 0 || userWidth <= 0 {
		return 0, errors.New("Не верно заданы параметры роста и веса")
	}

	return userWidth / math.Pow(userHeigth/100, 2), nil
}

func getUserInput() (float64, float64) {
	var userHeigth, userWidth float64
	fmt.Print("Внесите свой рост в сантиметрах: ")
	fmt.Scan(&userHeigth)
	fmt.Print("Внесите свой вес в килограммах: ")
	fmt.Scan(&userWidth)
	return userHeigth, userWidth
}

func userResponseCalculate() bool {
	var again string
	fmt.Println("Считаем еще, введите Y/N")
	fmt.Scan(&again)
	if again == "Y" || again == "y" {
		return true
	}
	return false
}

func test() {
	var y int
	fmt.Scan(&y)

	if y < 0 {
		panic(fmt.Sprintf("У Вас отрицательное число: %v", y))
	}

	for i := 1; i < 6; i++ {
		fmt.Println(i)
	}

	for {
		fmt.Print("Привет")
	}
}
