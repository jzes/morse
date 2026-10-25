package decoder_test

import (
	"bytes"
	"errors"
	"io"
	"strings"
	"testing"
	"testing/iotest"

	"github.com/jzes/morse/alphabet"
	"github.com/jzes/morse/decoder"
)

func TestDecode(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "empty input", input: "", want: "\n"},
		{name: "single dot", input: ".", want: "e\n"},
		{name: "single dash", input: "-", want: "t\n"},
		{name: "sos", input: "... --- ...", want: "sos\n"},
		{name: "word separator", input: ".... . .-.. .-.. ---  .-- --- .-. .-.. -..", want: "hello world\n"},
		{name: "multiple words", input: "--. ---  .. ...  ..-. ..- -.", want: "go is fun\n"},
		{name: "trailing letter separator", input: "... --- ... ", want: "sos\n"},
		{name: "trailing newline", input: "... --- ...\n", want: "sos\n"},
		{
			name:  "full alphabet",
			input: ".- -... -.-. -.. . ..-. --. .... .. .--- -.- .-.. -- -. --- .--. --.- .-. ... - ..- ...- .-- -..- -.-- --..",
			want:  "abcdefghijklmnopqrstuvwxyz\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var output bytes.Buffer
			d := decoder.New(&alphabet.RootNode)

			if err := d.Decode(strings.NewReader(tt.input), &output); err != nil {
				t.Fatalf("Decode() error = %v", err)
			}
			if got := output.String(); got != tt.want {
				t.Errorf("Decode(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestDecodeCustomTree(t *testing.T) {
	root := decoder.CodeNode{
		Dot:  &decoder.CodeNode{Key: "0"},
		Dash: &decoder.CodeNode{Key: "1"},
	}
	var output bytes.Buffer
	d := decoder.New(&root)

	if err := d.Decode(strings.NewReader(". - - ."), &output); err != nil {
		t.Fatalf("Decode() error = %v", err)
	}
	if got, want := output.String(), "0110\n"; got != want {
		t.Errorf("Decode() = %q, want %q", got, want)
	}
}

func TestDecodeReadError(t *testing.T) {
	wantErr := errors.New("reader failed")
	tests := []struct {
		name  string
		input io.Reader
	}{
		{name: "immediate failure", input: iotest.ErrReader(wantErr)},
		{
			name:  "failure after partial input",
			input: io.MultiReader(strings.NewReader("... "), iotest.ErrReader(wantErr)),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var output bytes.Buffer
			d := decoder.New(&alphabet.RootNode)

			err := d.Decode(tt.input, &output)
			if !errors.Is(err, wantErr) {
				t.Fatalf("Decode() error = %v, want wrapped %v", err, wantErr)
			}
			if output.Len() != 0 {
				t.Errorf("Decode() wrote %q despite read failure", output.String())
			}
		})
	}
}
