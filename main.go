package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/jzes/morse/alphabet"
)

func main() {

	mCode, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	bufferNode := &alphabet.RootNode
	itsASpace := false
	var word strings.Builder
	for _, b := range mCode {
		switch b {
		case 45:
			if bufferNode.Dash == nil {
				panic("invalid morse code detected")
			}
			bufferNode = bufferNode.Dash
			itsASpace = false
		case 46:
			if bufferNode.Dot == nil {
				panic("invalid morse code detected")
			}
			bufferNode = bufferNode.Dot
			itsASpace = false
		case 32:
			if itsASpace {
				word.WriteString(" ")
				break
			}
			itsASpace = true
			word.WriteString(bufferNode.Key)
			bufferNode = &alphabet.RootNode
		}
	}
	if bufferNode != &alphabet.RootNode {
		word.WriteString(bufferNode.Key)
	}

	fmt.Println("->", word.String())
	fmt.Println(string(mCode))
}
