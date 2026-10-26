package account

import (
	"demo/files"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (vault *Vault) FindAccountsByUrl(url string) []Account {
	var accounts []Account
	for _, acc := range vault.Accounts {
		if strings.Contains(acc.Url, url) {
			accounts = append(accounts, acc)
		}
	}
	return accounts
}

func (vault *Vault) DeleteAccountsByUrl(url string) []Account {
	var deleted, kept []Account
	for _, acc := range vault.Accounts {
		if strings.Contains(acc.Url, url) {
			deleted = append(deleted, acc)
			continue
		}
		kept = append(kept, acc)
	}
	vault.Accounts = kept

	data, err := vault.ToBytes()
	if err != nil {
		fmt.Println(err)
	}
	files.WriteFile(data, "data.json")

	return deleted 
}

func NewValut() *Vault {
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
		fmt.Println(err)
		return nil
	}
	return &vault
}

func (vault *Vault) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.UpdatedAt = time.Now()
	data, err := vault.ToBytes()
	if err != nil {
		fmt.Println(err)
		return
	}
	files.WriteFile(data, "data.json")
}

func (acc *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(acc)
	if err != nil {
		fmt.Println(err)
	}
	return file, nil
}
