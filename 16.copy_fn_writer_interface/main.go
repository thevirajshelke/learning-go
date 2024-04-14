package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	/*
		Copy function internally reads bytes from source in small chunks and puts it in the
		dest till all the data is read fully.

		Copy function expects the first param which is of type Writer & second param is of type Reader

		type Writer interface {
			Write(p []byte) (n int, err error)
		}

		type Reader interface {
			Read(p []byte) (n int, err error)
		}

		os.Stdout is of File type which internally uses the interface writer
	*/
	io.Copy(os.Stdout, resp.Body)
}
