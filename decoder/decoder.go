package decoder

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

type CodeNode struct {
	Key  string
	Dot  *CodeNode
	Dash *CodeNode
}

var invalidMorseError = errors.New("invalid morse code detected")

type Decoder struct {
	bufferNode *CodeNode
	rootNode   *CodeNode
}

func New(rootNode *CodeNode) *Decoder {
	return &Decoder{
		bufferNode: rootNode,
		rootNode:   rootNode,
	}
}

func (d *Decoder) Decode(reader io.Reader, writer io.Writer) error {
	rawCode, err := io.ReadAll(reader)
	if err != nil {
		return fmt.Errorf("failed to read input: %w", err)
	}

	decodedCode, err := d.decode(rawCode)
	if err != nil {
		return fmt.Errorf("failed to decode code: %w", err)
	}

	fmt.Fprintln(writer, decodedCode)
	return nil
}

func (d *Decoder) decode(rawCode []byte) (string, error) {
	var decodedCode strings.Builder
	itsSpace := false

	for _, b := range rawCode {
		switch b {
		case 45:
			d.setNode(d.bufferNode.Dash)
			itsSpace = false
		case 46:
			d.setNode(d.bufferNode.Dot)
			itsSpace = false
		case 32:
			if itsSpace {
				decodedCode.Write([]byte{32})
				break
			}
			itsSpace = true
			decodedCode.WriteString(d.bufferNode.Key)
			d.bufferNode = d.rootNode
		}
	}
	if d.bufferNode != d.rootNode {
		decodedCode.WriteString(d.bufferNode.Key)
	}
	return decodedCode.String(), nil
}

func (d *Decoder) setNode(n *CodeNode) error {
	if n == nil {
		return invalidMorseError
	}
	d.bufferNode = n
	return nil
}
