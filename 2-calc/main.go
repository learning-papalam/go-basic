package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var result float64

	operation := userOperation()
	data := userInput()

	switch operation {
	case "AVG":
		result = avg(data)
	case "SUM":
		result = sum(data)
	case "MED":
		result = med(data)
	default:
		panic("Неизвестная операция")
	}

	fmt.Printf("Результат операции %v будет %.2f", operation, result)
}

func userOperation() string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Введите операцию (AVG, SUM, MED): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		input = strings.ToUpper(input)
		switch input {
		case "AVG", "SUM", "MED":
			return input
		default:
			fmt.Println("Вы ввели не что-то непонятное!")
		}
	}
}

func userInput() []int {
	var nums []int
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите числа через запятую: ")
	data, _ := reader.ReadString('\n')
	data = strings.TrimSpace(data)
	data = strings.ReplaceAll(data, " ", "")
	parts := strings.Split(data, ",")

	for _, value := range parts {
		num, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println("Вы ввели не число!")
			continue
		}
		nums = append(nums, num)
	}
	return nums
}

func sum(nums []int) float64 {
	var sum int
	for _, value := range nums {
		sum += value
	}
	return float64(sum)
}

func avg(nums []int) float64 {
	sum := sum(nums)
	return sum / float64(len(nums))
}

func med(nums []int) float64 {
	if len(nums) == 0 {
		return 0
	}
	if len(nums)%2 == 0 {
		return float64(nums[len(nums)/2-1]+nums[len(nums)/2]) / 2.0
	}
	return float64(nums[len(nums)/2])
}
