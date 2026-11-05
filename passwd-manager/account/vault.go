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

type VaultWithDb struct {
	Vault
	db *files.JsonDb
}

func NewValut(db *files.JsonDb) *VaultWithDb {
	data, err := db.Read()
	if err != nil {
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db: db,
		}
	}
	var vault VaultWithDb
	err = json.Unmarshal(data, &vault)
	if err != nil {
		fmt.Println(err)
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db: db,
		}

	}
	vault.db = db
	return &vault
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

func (vault *VaultWithDb) DeleteAccountsByUrl(url string) []Account {
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

	vault.db.Write(data)
	return deleted
}

func (vault *VaultWithDb) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.save()
}

func (vault *VaultWithDb) save(){
	vault.UpdatedAt = time.Now()
	data, err := vault.ToBytes()
	if err != nil {
		fmt.Println(err)
		return
	}
	vault.db.Write(data)
}


func (acc *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(acc)
	if err != nil {
		fmt.Println(err)
	}
	return file, nil
}
