package main

import (
	"fmt"
	"os"
	"os/exec"
)



// displayHelp informs the user about our hardcoded functions
func displayHelp() {
    fmt.Printf(
        "Welcome to %v! These are the available commands: \n",
        cliName,
    )
    fmt.Println("help    - Show available commands")
	fmt.Println("map     - Show map of pokemon world")
    fmt.Println("clear   - Clear the terminal screen")
    fmt.Println("exit    - Closes your connection to ", cliName)
}
 
// clearScreen clears the terminal screen
func clearScreen() {
    cmd := exec.Command("clear")
    cmd.Stdout = os.Stdout
    cmd.Run()
}


