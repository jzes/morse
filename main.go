package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/jzes/morse/alphabet"
	"github.com/jzes/morse/decoder"
)

func main() {
	var reader io.Reader = os.Stdin
	if len(os.Args) > 1 {
		reader = strings.NewReader(os.Args[1])
	}

	decoder := decoder.New(&alphabet.RootNode)
	err := decoder.Decode(reader, os.Stdout)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
