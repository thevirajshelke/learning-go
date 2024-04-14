package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	logger := logWriter{}
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(logger, resp.Body)
}

func (logWriter) Write(p []byte) (int, error) {
	return fmt.Print(string(p))
}
