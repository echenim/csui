package main

import (
	"fmt"

	"github.com/echenim/csui/widgets"
)

func main() {
	//envname := []string{"Dev", "Prod", "Stage"}
	oldID, newID, err := widgets.MigrationComponent()
	if err == widgets.ErrInputCancelled {
		fmt.Println("User cancelled.")
		return
	}

	fmt.Printf("Old : %s\n New : %s \n", oldID, newID)
}
