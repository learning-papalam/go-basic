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
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updateAt"`
}

func (a *Account) generatePassword(n int) {
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	digits := "0123456789"
	specials := "~=+%^*/()[]{}/!@#$?|"

	all_symbols := fmt.Sprintf("%v%v%v", chars, digits, specials)
	all_rune := []rune(all_symbols)

	password := make([]rune, n)

	for i := range password {
		number := rand.IntN(len(all_rune))
		password[i] = all_rune[number]
	}

	a.Password = string(password)

}

func NewAccount(login, password, urlString string) (*Account, error) {
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}

	if _, err := url.ParseRequestURI(urlString); err != nil {
		return nil, errors.New("INVALID_URL")
	}

	newAcc := &Account{
		Login:     login,
		Password:  password,
		Url:       urlString,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if password == "" {
		newAcc.generatePassword(16)
	}

	return newAcc, nil

}

func (a *Account) OutputAccaunt() {
	color.Cyan(a.Login)
	color.Green(a.Password)
	color.Green(a.Url)
}
