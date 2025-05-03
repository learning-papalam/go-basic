package main

import (
	"fmt"
	"strings"
)

func main() {
	// russians := "Привет мир"
	// fmt.Printf("Начинается ли слово %v \n", isSearchSub(russians, "Пр"))
	// fmt.Println(strings.Replace("Привет мир ППППП", "П", "п", -1))
	// fmt.Println("Ищем вхождение: ", strings.Contains("Привет мир", "мир"))
	// fmt.Println(strings.ToUpper("привет мир"))
	// fmt.Println(join("Доброй", "ночи"))
	// a := []string{"Доброе", "утро"}
	// fmt.Println(strings.Join(a, " "))
	pallindrom := "Топот"
	fmt.Printf("Это паллиндром: %v %v", pallindrom, isPalindrom(pallindrom))

}

func isSearchSub(str string, sub string) bool {
	_, found := strings.CutPrefix(str, sub)
	return found
}

func join(str1 string, str2 string) string {
	return str1 + " " + str2
}

func isPalindrom(str string) bool {
	normalizeStr := strings.ToLower(str)

	var result string

	for _, v := range normalizeStr {
		result = string(v) + result
	}

	if normalizeStr == result {
		return true
	}
	return false
}
