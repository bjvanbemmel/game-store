package cmd

import (
	"io"
	"net/http"
	"testing"
	"time"
)

func TestServe(t *testing.T) {
	var expect string = "Pong!"

	serve := Serve{}
	go serve.Execute()

	var resp *http.Response
	var err error

	for i := 0; i < 5; i++ {
		resp, err = http.Get("http://127.0.0.1:8123/api/ping")
		if err != nil {
			time.Sleep(time.Second)

			continue
		}
		defer resp.Body.Close()
	}

	if err != nil {
		t.Fatal(err.Error())
	}

	response, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err.Error())
	}

	if string(response) != expect {
		t.Fatalf("Expected %s, got %s.\n%v", expect, response, err)
	}
}
