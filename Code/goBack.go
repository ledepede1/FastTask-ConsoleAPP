package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/TwiN/go-color"
)

func goBack() {
	fmt.Print("\n")
	fmt.Println(color.InBlackOverYellow("Go back to main menu? (yes): "))
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	goBack := scanner.Text()

	if strings.EqualFold("yes", goBack) {
		fmt.Print("\n")
		main()
	}
	fmt.Print("\n")
}
