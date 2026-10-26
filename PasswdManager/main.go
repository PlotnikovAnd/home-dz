package main

import (
	"demo/account"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	welcomeMessage := "Меню:\n" +
		`1 - Создать акк, 2 - Найти акк, 3 - Удалить акк, 4 - Вывести все, Другое - Выход` + "\n"

	vault := account.NewValut()

Menu:
	for {
		in := promptData(welcomeMessage)

		switch in {
		case "1":
			createAccount(vault)
		case "2":
			findAccount(vault)
		case "3":
			deleteAccount(vault)
		case "4":
			printAccounts(vault)
		default:
			break Menu
		}
	}
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scanln(&res)
	return res
}

func createAccount(vault *account.Vault) {
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	url := promptData("Введите URL: ")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		color.Red(string(err.Error()))
		return
	}
	vault.AddAccount(*myAccount)
}

func findAccount(vault *account.Vault) {
	url := promptData("Введите сайт для поиска: ")
	accounts := vault.FindAccountsByUrl(url)
	if len(accounts) == 0 {
		color.Red("Аккаунты не найдены")
		return
	}
	for _, acc := range accounts {
		acc.Output()
	}
}

func deleteAccount(vault *account.Vault) {
	// URL
	// метод удаления vault
	// проверка удалено или нет
	url := promptData("Введите сайт для удаления: ")
	accs := vault.DeleteAccountsByUrl(url)
	if len(accs) == 0 {
		color.Red("Аккаунты не найдены")
		return
	}
	resString := fmt.Sprintln("Аккаунты удалены:\n" , accs)
	color.Green(resString)
}

func printAccounts(vault *account.Vault){
	for _, acc := range vault.Accounts {
		color.Green(fmt.Sprint(acc.Url, " " , acc.CreatedAt))
	}
}
