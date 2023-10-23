package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/TwiN/go-color"
	"github.com/inancgumus/screen"
)

func main() {
	fmt.Println(color.InBlackOverBlue("Fast Task"))
	printOptions()

	fmt.Print(color.InBlackOverYellow("$ Input:"), " ")
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	choosenOption(input)
}

func choosenOption(input string) {
	if input != "1" && input != "2" && input != "3" { // Pressed enter
		if strings.ToLower(input) == "close" {
			os.Exit(0)
		}
		main()
	}

	if input == "1" {
		screen.Clear()
		writeTask()
	}
	if input == "2" {
		screen.Clear()
		readTask()
	}
	if input == "3" {
		screen.Clear()
		editTask()
	}
}

func printOptions() {
	fmt.Printf(`
Type 1 for writing a new Task!
Type 2 for reading a Task!
Type Close for closing the application!
`)
}
