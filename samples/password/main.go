package main

import (
	"fmt"

	"github.com/echenim/csui/widgets"
)

func main() {

	password, err := widgets.PasswordInput("Enter your password...")
	if err == widgets.ErrInputCancelled {
		fmt.Println("User cancelled.")
		return
	}

	fmt.Printf("Your password is %s!\n", password)
}
