package cmd

import (
	"errors"
	"testing"
)

func TestCommandShouldSucceed(t *testing.T) {
	var handler Handler = Handler{}

	for _, arg := range []string{"help", "serve", "migrate", "seed"} {
		_, err := handler.Match(arg)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestCommandShouldFail(t *testing.T) {
	var handler Handler = Handler{}

	_, err := handler.Match("N/A")
	if !errors.Is(err, ArgNotFoundErr) {
		t.Fatalf("Expected %q error, got %v.", ArgNotFoundErr, err)

		return
	}
}
