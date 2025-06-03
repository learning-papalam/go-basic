package main

import (
	"demo/password/account"
	"demo/password/encrypter"
	"demo/password/files"
	"demo/password/output"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

var menu = map[string]func(*account.VaultWithDB){
	"1": createAccount,
	"2": searchAccountByURL,
	"3": searchAccountByLogin,
	"4": deleteAccount,
}

var menuVariants = []string{
	"1. Создать аккаунт [1]",
	"2. Найти аккаунт по URL [2]",
	"3. Найти аккаунт по логину [3]",
	"4. Удалить аккаунт [3]",
	"5. Выход [любая клавиша]",
	"Выбрите пункт меню",
}

func main() {

	err := godotenv.Load()
	if err != nil {
		output.PrintError("Не найден файл .env")
	}

	vault := account.NewVault(files.NewJsonDb("vault.db"), *encrypter.NewEncrypter())

	for {
		variant := promtUser(menuVariants...)
		menuFunc := menu[variant]
		if menuFunc == nil {
			break
		}
		menuFunc(vault)
	}

}

func searchAccountByURL(vault *account.VaultWithDB) {
	var url string = promtUser("Введите URL для поиска")

	foundAcc := vault.FindAccounts(url, func(acc account.Account, url string) bool {
		urlNormalazed := strings.ToLower(url)
		return strings.Contains(strings.ToLower(acc.Url), urlNormalazed)
	})
	outputResult(&foundAcc)
}

func searchAccountByLogin(vault *account.VaultWithDB) {
	var login string = promtUser("Введите логин для поиска")

	foundAcc := vault.FindAccounts(login, func(acc account.Account, login string) bool {
		loginNormalazed := strings.ToLower(login)
		return strings.Contains(strings.ToLower(acc.Login), loginNormalazed)
	})
	outputResult(&foundAcc)
}

func outputResult(foundAcc *[]account.Account) {
	if len(*foundAcc) == 0 {
		output.PrintError("Аккаунтов не найдено")
	}
	for _, account := range *foundAcc {
		account.OutputAccaunt()
	}
}

func deleteAccount(vault *account.VaultWithDB) {
	var url string = promtUser("Введите URL для удаления")
	isDeleted := vault.DeleteAccountByURL(url)
	if isDeleted {
		color.Green("Удалено")
		return
	}
	output.PrintError("Аккаунт не найден")
}

func createAccount(vault *account.VaultWithDB) {
	var login string = promtUser("Введите логин")
	var password string = promtUser("Введите пароль")
	var url string = promtUser("Введите URL")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		panic(err.Error())
	}
	myAccount.OutputAccaunt()

	vault.AddAccount(myAccount)
}

func promtUser(promt ...string) string {
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
