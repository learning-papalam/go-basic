package account

import (
	"demo/password/output"
	"encoding/json"
	"strings"
	"time"
)

type Db interface {
	Read() ([]byte, error)
	Write([]byte)
}

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type VaultWithDB struct {
	Vault
	db Db
}

func NewVault(db Db) *VaultWithDB {
	data, err := db.Read()
	if err != nil {
		return &VaultWithDB{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db: db,
		}
	}

	var vault Vault

	err = json.Unmarshal(data, &vault)
	if err != nil {
		output.PrintError("Не удалось прочитать JSON")
	}
	return &VaultWithDB{
		Vault: vault,
		db:    db,
	}
}

func (v *Vault) ToByte() ([]byte, error) {
	data, err := json.Marshal(v)
	return data, err
}

func (v *VaultWithDB) AddAccount(a *Account) {
	v.Accounts = append(v.Accounts, *a)
	v.save()
}

func (v *VaultWithDB) FindAccountByURL(url string) []Account {
	var foundAcc []Account
	urlNormalazed := strings.ToLower(url)
	for _, acc := range v.Accounts {
		if strings.Contains(strings.ToLower(acc.Url), urlNormalazed) {
			foundAcc = append(foundAcc, acc)
		}
	}
	return foundAcc
}

func (v *VaultWithDB) DeleteAccountByURL(url string) bool {
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

func (v *VaultWithDB) save() {
	v.UpdatedAt = time.Now()

	data, err := v.Vault.ToByte()
	if err != nil {
		output.PrintError("Не удалось преобразовать JSON")
	}
	v.db.Write(data)
}
