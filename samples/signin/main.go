package main

import (
	"fmt"

	"github.com/echenim/csui/widgets"
)

func main() {
	envname := []string{"Dev", "Prod", "Stage"}
	env, username, password, err := widgets.SignInInput("Select Enviromemnt...", "Enter your name...", "Enter your password...", envname)
	if err == widgets.ErrInputCancelled {
		fmt.Println("User cancelled.")
		return
	}

	fmt.Printf("Enviroment : %s\nYour username : %s \npassword : %s!\n", env, username, password)
}
