package conio

import (
	"os"
	"os/exec"
	"runtime"
)

func clearScreen() {
	var clearCmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		clearCmd = exec.Command("cmd", "/c", "cls")
	default:
		clearCmd = exec.Command("clear")
	}
	clearCmd.Stdout = os.Stdout
	clearCmd.Run()
}
