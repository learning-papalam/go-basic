package main

import "fmt"

type arr = [4]int

func main() {
	a := [4]int{1, 2, 3, 4}
	revers(&a)
	fmt.Println(a)

	var pointer *int

	fmt.Println(*pointer)

	var p *string // Это nil-указатель
	if p == nil {
		fmt.Println("Указатель равен nil")
	}


}

func revers(a *arr) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
	}
}
