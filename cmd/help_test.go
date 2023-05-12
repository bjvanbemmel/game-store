package cmd

import (
	"bytes"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/bjvanbemmel/game-store/docs"
)

var embedEmptyErr error = errors.New("docs/embed.txt is empty")

func TestHelp(t *testing.T) {
	if len(docs.HelpEmbed) == 0 {
		t.Fatal(embedEmptyErr)
	}

	var expect string = string(docs.HelpEmbed)
	var help Help = Help{}

	origStdOut := os.Stdout
	r, w, err := os.Pipe()

	if err != nil {
		t.Fatal(err)
	}

	os.Stdout = w

	help.Execute()

	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	err = w.Close()
	os.Stdout = origStdOut
	out := <-outC

	if out != string(docs.HelpEmbed) {
		t.Fatalf("Expected %s, got %s.\n%v", expect, out, err)
	}
}
