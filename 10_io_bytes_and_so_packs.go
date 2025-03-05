// example code of io and bytes package

package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	//Write the content of the buffer to the standard output
	defer func() {
		if _, err := io.WriteString(os.Stdout, "Hello from io package"); err != nil {
			fmt.Println(err)
		}
	}()

	//Create a new buffer value and write a string to it
	var b bytes.Buffer
	b.Write([]byte("Hello"))

	//Use the Fprintf function to concatenate a string to the buffer
	fmt.Printf("%s from fmt package with bytes\n", b.String())
}
