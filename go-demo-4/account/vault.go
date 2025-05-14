package account

import (
	"demo/password/files"
	"encoding/json"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewVault() *Vault {
	data, err := files.ReadFile("data.json")
	if err != nil {
		return &Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}
	}

	var vault Vault

	err = json.Unmarshal(data, &vault)
	if err != nil {
		color.Red("Не удалось прочитать JSON")
	}
	return &vault
}

func (v *Vault) ToByte() ([]byte, error) {
	data, err := json.Marshal(v)
	return data, err
}

func (v *Vault) AddAccount(a *Account) {
	v.Accounts = append(v.Accounts, *a)
	v.save()
}

func (v *Vault) FindAccountByURL(url string) []Account {
	var foundAcc []Account
	urlNormalazed := strings.ToLower(url)
	for _, acc := range v.Accounts {
		if strings.Contains(strings.ToLower(acc.Url), urlNormalazed) {
			foundAcc = append(foundAcc, acc)
		}
	}
	return foundAcc
}

func (v *Vault) DeleteAccountByURL(url string) bool {
	isDeleted := false
	urlNormalazed := strings.ToLower(url)
	i := 0
	for _, account := range v.Accounts {
		if !strings.Contains(strings.ToLower(account.Url), urlNormalazed) {
			v.Accounts[i] = account
			i++
		} else {
			isDeleted = true
		}
	}

	v.Accounts = v.Accounts[:i]
	v.save()

	return isDeleted
}

func (v *Vault) save() {
	v.UpdatedAt = time.Now()

	data, err := v.ToByte()
	if err != nil {
		color.Red("Не удалось преобразовать JSON")
	}
	files.WriteFile(data, "data.json")
}
