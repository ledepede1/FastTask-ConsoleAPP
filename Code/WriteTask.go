package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/TwiN/go-color"
)

var path = "D:/Text/"

var taskTitle string
var taskText []string

func writeTask() {
	fmt.Print("\n")
	fmt.Println(color.InBlackOverBlue("$ Task"))
	fmt.Print("\n")

	fmt.Println(color.InBlackOverYellow("Title: "))
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	taskTitle = scanner.Text()
	fmt.Print("\n")

	fmt.Println(color.InBlackOverYellow("Task (Two spaces for saving): "))
	reader := bufio.NewReader(os.Stdin)

	var lines []string
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		if len(strings.TrimSpace(line)) == 0 {
			break
		}
		lines = append(lines, line)
	}

	taskText = append(taskText, lines...)

	saveTask(taskTitle, taskText)
}

func saveTask(tile string, text []string) {
	createFile, err := os.Create(path + taskTitle + ".txt")
	if err != nil {
		panic(err)
	}

	defer createFile.Close()

	createFile.WriteString(strings.Join(text, ""))
	fmt.Println("Wrote file to", createFile.Name())

	goBack()
}
