package screen

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func ClearScreen() {
	cmd := exec.Command("clear")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Menu() {
	var input string
	for {
		fmt.Printf("\tMenu\n")
		fmt.Printf("\t1. Register client\n")
		fmt.Printf("\tS. Exit\n")
		fmt.Printf("\t: ")
		if _, err := fmt.Scanln(&input); err != nil {
			fmt.Println("\tInvalid option")
			return
		}
		if strings.TrimSpace(input) == "s" {
			break
		}
		ClearScreen()
	}
}
