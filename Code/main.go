package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/TwiN/go-color"
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
	if input != "1" && input != "2" { // Pressed enter
		if input == "close" {
			os.Exit(0)
		}
		main()
	}

	if input == "1" {
		writeTask()
	}
}

func printOptions() {
	fmt.Printf(` 
Type 1 for writing a new Task!
Type Close for closing the application!
`)
}
