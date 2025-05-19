package main

import (
	"demo/password/account"
	"demo/password/files"
	"demo/password/output"
	"fmt"

	"github.com/fatih/color"
)

func main() {

	vault := account.NewVault(files.NewJsonDb("data.json"))

Loop:
	for {
		variant := promtUser([]string{
			"1. Создать аккаунт [1]",
			"2. Найти аккаунт [2]",
			"3. Удалить аккаунт [3]",
			"4. Выход [любая клавиша]",
			"Выбрите пункт меню",
		})
		switch variant {
		case "1":
			createAccount(vault)
		case "2":
			searchAccount(vault)
		case "3":
			deleteAccount(vault)
		default:
			fmt.Println("Выход из программы")
			break Loop
		}
	}

}

func searchAccount(vault *account.VaultWithDB) {
	var url string = promtUser([]string{"Введите URL для поиска"})

	foundAcc := vault.FindAccountByURL(url)
	if len(foundAcc) == 0 {
		output.PrintError("Аккаунтов не найдено")
	}
	for _, account := range foundAcc {
		account.OutputAccaunt()
	}
}

func deleteAccount(vault *account.VaultWithDB) {
	var url string = promtUser([]string{"Введите URL для удаления"})
	isDeleted := vault.DeleteAccountByURL(url)
	if isDeleted {
		color.Green("Удалено")
		return
	}
	output.PrintError("Аккаунт не найден")
}

func createAccount(vault *account.VaultWithDB) {
	var login string = promtUser([]string{"Введите логин"})
	var password string = promtUser([]string{"Введите пароль"})
	var url string = promtUser([]string{"Введите URL"})

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		panic(err.Error())
	}
	myAccount.OutputAccaunt()

	vault.AddAccount(myAccount)
}

func promtUser[T any](promt []T) string {
	var userData string

	for i, line := range promt {
		if i == len(promt)-1 {
			fmt.Printf("%v: ", line)
		} else {
			fmt.Println(line)
		}
	}

	fmt.Scanln(&userData)
	return userData
}
