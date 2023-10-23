package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/TwiN/go-color"
)

func readTask() {
	fmt.Print("\n")
	fmt.Println(color.InBlackOverBlue("$ Task"))
	fmt.Print("\n")

	fmt.Println("Type in the path to the task: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	pathToRead := scanner.Text()

	file, err := os.ReadFile(pathToRead)
	if err != nil {
		panic(err)
	}
	fileStr := string(file)
	fmt.Println() // Blank space!
	fmt.Print(color.InBlackOverYellow("$ " + string(pathToRead)))
	fmt.Println() // Blank space!
	fmt.Println() // Blank space!
	fmt.Println(fileStr)

	goBack()
}
