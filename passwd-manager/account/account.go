package account

import (
	"errors"
	"math/rand/v2"
	"net/url"
	"time"

	"strings"

	"github.com/fatih/color"
)

type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (acc *Account) Output() {
	res := acc.Login + "\t " +acc.Password +"\t "+ acc.Url
	color.Cyan(res)
}

func (acc *Account) generatePassword(n int) {
	// range [33-127)
	var result strings.Builder
	for range n {
		ch := rune(rand.IntN(127-33) + 33)
		result.WriteString(string(ch))
	}
	acc.Password = result.String()
}

func NewAccount(login, password, urlString string) (*Account, error) {

	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}

	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}

	newAcc := &Account{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Login:     login,
		Password:  password,
		Url:       urlString,
	}
	if password == "" {
		newAcc.generatePassword(10)
	}

	return newAcc, nil
}
