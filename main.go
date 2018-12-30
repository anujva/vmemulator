package main

import (
	"context"
	"fmt"
	"os"

	"github.com/anujva/vmemulator/parser"
)

func main() {
	// This will need to read a file, iterate over the lines
	// and then process it as needed. The file iterator implementation can be borrowed from
	// the last assembler project that we did.
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Please supply the text file that will be used to translate")
		os.Exit(1)
	}

	ctx := context.Background()
	f, _ := os.Open(args[0])
	parser := parser.New(f)
	for parser.HasMoreCommands() {
		commands, err := parser.Parse(ctx)
		if err != nil {
			fmt.Println("Error occurred", err)
		}
		fmt.Println(commands)
	}
}
