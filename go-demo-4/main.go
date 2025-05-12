package main

import (
	"demo/password/account"
	"demo/password/files"
	"fmt"
)

func main() {
	var login string = promtUser("Введите логин")
	var password string = promtUser("Введите пароль")
	var url string = promtUser("Введите URL")

	account1, err := account.NewAccountWithTimeStamp(login, password, url)
	if err != nil {
		panic(err.Error())
	}
	files.WriteFile("Привет, мир!!!", "text.txt")
	account1.OutputAccaunt()

}

func promtUser(promt string) string {
	var userData string
	fmt.Printf("%v: ", promt)

	fmt.Scanln(&userData)

	return userData
}
