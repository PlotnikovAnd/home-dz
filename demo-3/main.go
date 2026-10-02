package main

import (
	"fmt"
	
	"github.com/fatih/color"
	"demo/file"
	"demo/account"
)


func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scanln(&res)
	return res
}



func main() {
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	url := promptData("Введите URL: ")

	acc1, err := account.NewAccount(
		login,
		password,
		url,
	)
	file.ReadFile()
	if err != nil {
		color.Red(string(err.Error()))
		return
	}

	acc1.OutputAccount()
}
