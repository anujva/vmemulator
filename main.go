package main

import (
	"context"
	"fmt"
	"os"

	"github.com/anujva/iterator/fileiterator"
)

func main() {
	// This will need to read a file, iterate over the lines
	// and then process it as needed. The file iterator implementation can be borrowed from
	// the last assembler project that we did.
	f, _ := os.Open("/Users/anujvarma/code/nand2tetris/projects/07/VMAbstraction.txt")
	fileIterator := fileiterator.New(f)
	ctx := context.Background()

	for fileIterator.HasNext() {
		(fileIterator.Next(ctx))
		//fmt.Println(*strPtr)
	}

	fileIterator.Reset()
	for fileIterator.HasNext() {
		strPtr := (fileIterator.Next(ctx)).(*string)
		fmt.Println(*strPtr)
	}
}
