package main

import "fmt"

type stringMap map[string]string

func main() {
	m := make(stringMap, 5)
Menu:
	for {
		fmt.Println("Enter cmd (1 = show, 2 = add, 3 = delete, 4 = exit):")
		cmd := 0
		fmt.Scan(&cmd)

		switch cmd {
		case 1:
			for key, value := range m {
				fmt.Printf("key:%s, value:%s\n", key, value)
			}
		case 2:
			var key, value string
			fmt.Print("Enter key and value to add: ")
			fmt.Scan(&key, &value)
			m[key] = value
		case 3:
			var key string
			fmt.Print("Enter key to delete: ")
			fmt.Scan(&key)
			delete(m, key)
			fmt.Printf("deleted %s key\n", key)
		case 4:
			break Menu
		default:
			continue
		}
	}
}
