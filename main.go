package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/adityatresnobudi/bin2dec/cli"
)

func promptInput(scanner *bufio.Scanner, text string) string {
	fmt.Print(text)
	scanner.Scan()
	return scanner.Text()
}

func outputHandler(err error, a ...interface{}) {
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println()
		return
	}
	fmt.Println(a...)
}

func main() {
	exit := false
	scanner := bufio.NewScanner(os.Stdin)
	menu := "Binary Converter\n\n" +
		"1. Convert\n" +
		"2. Exit"
		
	for !exit {
		fmt.Println(menu)
		input := promptInput(scanner, "Enter Menu :")

		switch input {
		case "1":
			binary := promptInput(scanner, "Enter Binary number :")
			res, err := cli.HandleConvert(binary)
			outputHandler(err, res)
		case "2":
			exit = true
			fmt.Println("")
			fmt.Println("Good bye!")
		default:
			fmt.Println("")
			fmt.Println("Invalid Menu")
		}
	}
}
