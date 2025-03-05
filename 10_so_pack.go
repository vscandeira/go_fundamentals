// example code with so package to read the file 01_number.go as a text file
// and print it to the console
// the file 01_number.go should be in the same directory as this file
// and is already created

package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("01_number.go")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer file.Close()

	// read the file
	// create a byte slice to store the content of the file
	// the size of the byte slice is the size of the file
	// read the content of the file to the byte slice
	// print the content of the file
	// the content of the file is in byte slice format
	// convert the byte slice to string to print it
	fileInfo, _ := file.Stat()
	fileSize := fileInfo.Size()
	content := make([]byte, fileSize)
	file.Read(content)
	fmt.Println(string(content))
}
