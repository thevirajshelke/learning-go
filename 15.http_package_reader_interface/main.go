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

	// The following line prints a lot more than body and the body is not readily available
	// resp is a pointer to the response struct and reponse is a struct which contains many fields like Status, StatusCode. Proto etc.
	// One of the field in struct is Body which has a type of io.ReadCloser - ideally should be a blob of text, JSON etc.
	fmt.Println(resp)

	/*
		ReadCloser is a interface as shown,

		type ReadCloser interface {
			Reader
			Closer
		}

		Reader is also an interface

		type Reader interface {
			Read(p []byte) (n int, err error)
		}

		type Closer interface {
			Close() error
		}

		So basically,

		A struct can have a field which has no concrete type instead is a an interface. An interface can be created not only with functions but
		with other interfaces as well.

		Here what we are saying is that Body can have any type (value) as long as that type has a Read method and a Close method defined as receiver.
	*/

	/*
		This will create a slice with 0 space and read function cannot automatically resize. It will stop filling data
		when it sees that the byte slice is full. If you send an empty byte slice it thinks it's full and doesn't do anything
	*/
	// body := []byte{} // will read 0 bytes
	// body := make([]byte, 1) // will read 1 byte from the body

	fmt.Println("Response content length", resp.ContentLength)
	var body []byte
	if resp.ContentLength > 0 {
		body = make([]byte, resp.ContentLength)
	} else {
		body = make([]byte, 99999) // arbitary size which might result into loss of data (incomplete data filled)
	}
	bytesRead, err := resp.Body.Read(body) // Read is available on type io.ReadCloser and it accepts a byte slice that it will fill up
	if err != nil && err != io.EOF {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println(string(body))
	fmt.Println("No of bytes read", bytesRead)

}
