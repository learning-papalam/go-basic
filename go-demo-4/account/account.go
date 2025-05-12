package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color"
)

type Account struct {
	login    string
	password string
	url      string
}

func (a *Account) generatePassword(n int) {
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	digits := "0123456789"
	specials := "~=+%^*/()[]{}/!@#$?|"

	all_symbols := fmt.Sprintf("%v%v%v", chars, digits, specials)
	all_rune := []rune(all_symbols)

	password := make([]rune, n)

	for range n {
		number := rand.IntN(len(all_rune))
		password = append(password, all_rune[number])
	}

	a.password = string(password)

}

type AccountWithTimeStamp struct {
	createAt time.Time
	updateAt time.Time
	Account
}

func NewAccountWithTimeStamp(login, password, urlString string) (*AccountWithTimeStamp, error) {
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}

	if _, err := url.ParseRequestURI(urlString); err != nil {
		return nil, errors.New("INVALID_URL")
	}

	newAcc := &AccountWithTimeStamp{
		createAt: time.Now(),
		updateAt: time.Now(),
		Account: Account{
			login:    login,
			password: password,
			url:      urlString,
		},
	}

	if password == "" {
		newAcc.generatePassword(16)
	}

	return newAcc, nil

}

func (a *AccountWithTimeStamp) OutputAccaunt() {
	color.Cyan(a.login)
}
