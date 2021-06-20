package main

import (
	"bufio"
	"fmt"
	"os"
	"roman"
	"strconv"
	"strings"
)

func main() {

	for {

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Please enter some numbers or numerals, if you want to stop please enter 'quit':\n")
		stringValue, _ := reader.ReadString('\r')
		stringValue = strings.TrimSpace(stringValue)
		if strings.ToLower(stringValue) == "quit" {
			return
		}

		intValue, err := strconv.Atoi(strings.TrimSpace(stringValue))
		if err == nil {
			fmt.Println("Input interpreted as a integer, attempting to convert to numerals ...")
			roman.ToNumerals(intValue)
		} else {
			fmt.Println("Input interpreted as string, attempting to convert to numerals ...")
			roman.ToInt(stringValue)
		}
	}

}
