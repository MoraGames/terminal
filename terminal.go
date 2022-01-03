package terminal

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

//Pause prints a warning and waits for the input of any value or newline character
func Pause(onlyEnter bool) {
	if onlyEnter == true {
		fmt.Print("Press 'Enter' to continue...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	} else {
		fmt.Print("Press any key to continue...")
		bufio.NewReader(os.Stdin).ReadByte()
	}
}

//Clear runs the command to clean up the terminal window
func Clear() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "linux":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Println("[ERROR] Unable to clean your terminal window.")
	}
}
