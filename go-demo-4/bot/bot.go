package bot

// func sum(nums ...int) int {
// 	result := 0
// 	for _, v := range nums {
// 		result += v
// 	}
// 	return result
// }

// sum(1, 4, 6)

// type fn = func(int, int) int

// func calculate(a, b int, operation fn){
// 	result := operation(a, b)
// 	fmt.Printf("Результат %d", result)
// }

// calculate(5, 4, func(a, b int) int {
// 	return a + b
// })


// func counter() func() int {
// 	i := 0
// 	return func() int {
// 		i++
// 		return i
// 	}
// }

// fn := counter()
// fmt.Println(fn())  // 1
// fmt.Println(fn())  // 2
// fmt.Println(fn())  // 3
// fmt.Println(fn())  // 4

// func sum() {
// 	var fn = func(int, int) int {

// 	}
// }


type fn = func(int, int) int

func sun() fn {
	return func(a, b int) int {
		return a + b
	}
}