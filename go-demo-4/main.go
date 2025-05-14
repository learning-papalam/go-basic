package main

import (
	"demo/password/account"
	"demo/password/files"
	"fmt"
	"strings"

	"github.com/fatih/color"
)

func main() {

	vault := account.NewVault()

Loop:
	for {
		variant := getMenu()
		switch variant {
		case 1:
			createAccount(vault)
		case 2:
			searchAccount(vault)
		case 3:
			deleteAccount(vault)
		default:
			fmt.Println("Выход из программы")
			break Loop
		}
	}

}

func getMenu() int {
	var menuItem int
	menu := createMenu()

	fmt.Println("Выбрите пункт меню:")
	fmt.Printf("%v", strings.Join(*menu, "\n"))
	fmt.Print("\nВаш выбор: ")
	fmt.Scanln(&menuItem)

	return menuItem
}

func createMenu() *[]string {
	return &[]string{
		"1. Создать аккаунт [1]",
		"2. Найти аккаунт [2]",
		"3. Удалить аккаунт [3]",
		"4. Выход [любая клавиша]",
	}
}

func searchAccount(vault *account.Vault) {
	var url string = promtUser("Введите URL для поиска")

	foundAcc := vault.FindAccountByURL(url)
	if len(foundAcc) == 0 {
		color.Red("Аккаунтов не найдено")
	}
	for _, account := range foundAcc {
		account.OutputAccaunt()
	}
}

func deleteAccount(vault *account.Vault) {
	var url string = promtUser("Введите URL для удаления")
	isDeleted := vault.DeleteAccountByURL(url)
	if isDeleted {
		color.Green("Удалено")
		return
	}
	color.Red("Аккаунт не найден")
}

func createAccount(vault *account.Vault) {
	var login string = promtUser("Введите логин")
	var password string = promtUser("Введите пароль")
	var url string = promtUser("Введите URL")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		panic(err.Error())
	}
	myAccount.OutputAccaunt()

	vault.AddAccount(myAccount)

	data, err := vault.ToByte()
	if err != nil {
		panic(err.Error())
	}

	files.WriteFile(data, "data.json")
}

func promtUser(promt string) string {
	var userData string
	fmt.Printf("%v: ", promt)

	fmt.Scanln(&userData)

	return userData
}
