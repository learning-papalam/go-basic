package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	result := reverse(nums)
	fmt.Println(result)
}

// func doSomething(nums []int) []int {
// 	var result []int
// 	for i := len(nums); i > 0; i-- {
// 		result = append(result, nums[i-1])
// 	}
// 	return result
// }

func reverse(nums []int) []int {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}
