package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"strings"
)

type account struct {
	login    string
	password string
	url      string
}

func (acc *account) OutputAccount() {
	fmt.Println(&acc)
}

func (acc *account) generatePassword(n int) {
	// range [33-127)
	var result strings.Builder
	for range n {
		ch := rune(rand.IntN(127-33) + 33)
		result.WriteString(string(ch))
	}
	acc.password = result.String()
}

func NewAccount(login, password, urlString string) (*account, error) {

	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}

	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}

	newAcc := account{
		login: login,
		password: password,
		url: urlString,
	}

	if password == "" {
		newAcc.generatePassword(10)
	} 

	return &newAcc, nil
}

