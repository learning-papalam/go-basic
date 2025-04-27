package main

import "fmt"

func main() {
	const USDtoEUR float64 = 0.8796
	const USDtoRUB float64 = 82.65
	const UERtoRUB float64 = USDtoRUB / USDtoEUR

	fmt.Print(USDtoEUR, USDtoRUB, UERtoRUB)

}