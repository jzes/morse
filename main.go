package main

import (
	"os"

	"github.com/jzes/morse/alphabet"
	"github.com/jzes/morse/decoder"
)

func main() {
	decoder := decoder.New(&alphabet.RootNode)
	err := decoder.Decode(os.Stdin, os.Stdout)
	if err != nil {
		panic(err)
	}
}
