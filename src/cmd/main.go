package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"search"
)

func main() {
	for {
		fmt.Println("Please type input below and enter ? to start processing")

		// reads  all lines in one go unitl first occurence of delimiter '?'
		reader := bufio.NewReader(os.Stdin)
		inputString, err := reader.ReadString('?')
		if err != nil {
			fmt.Printf("Failed to read input. [Error:%s]\n", err)
		}

		inputString = strings.TrimSpace(inputString)
		lines := strings.Split(inputString[:len(inputString)-2], "\n") // excluding delimter from processesing
		for _, line := range lines {
			line = strings.TrimSpace(line)
			args := strings.Split(line, " ")
			search.Process(args) // process each cmd(P/Q)
		}
	}
}
