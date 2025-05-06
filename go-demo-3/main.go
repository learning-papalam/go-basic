package main

import (
	"errors"
	"fmt"
)

type bookmarksMap = map[string]string

func main() {
	var bookmarks = bookmarksMap{}
	showMenu()
	changeBookmark(bookmarks)
	fmt.Println(bookmarks)

Menu:
	for {
		operation := userOperation()
		switch operation {
		case 1:
			printAllBookmarks(bookmarks)
		case 2:
			key := getUserValue("ключ")
			value := getUserValue("значение")
			bookmarks[key] = value
		case 3:
			key := getUserValue("ключ, который нужно удалить")
			delete(bookmarks, key)
		case 4:
			fmt.Println("Выход из программы")
			break Menu
		default:
			panic(errors.New("неопознанная операция"))
		}

	}
}

func changeBookmark(m bookmarksMap) {
	m["Яндекс"] = "https://google.com"
	fmt.Println(m)

}

func showMenu() {
	fmt.Println("1. Посмотреть все закладки")
	fmt.Println("2. Добавить закладку")
	fmt.Println("3. Удалить закладку")
	fmt.Println("4. Выход")
}

func printAllBookmarks(m bookmarksMap) {
	for key, value := range m {
		fmt.Printf("Ключ: %v, значение: %v\n", key, value)
	}
}

func getUserValue(name string) string {
	var value string
	fmt.Printf("Напишите %v: ", name)
	fmt.Scan(&value)
	return value
}

func userOperation() int {
	var result int
	for {
		fmt.Print("Сделайте выбор (1, 2, 3 или 4): ")
		fmt.Scan(&result)
		switch result {
		case 1, 2, 3, 4:
			return result
		default:
			continue
		}
	}
}
