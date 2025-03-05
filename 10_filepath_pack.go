//example code with walk function from filepath package
//to read from the current directory and return everything found inside my path

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// walk function from filepath package
	// it will walk through the current directory
	// and return everything found inside my path
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}
		if strings.Contains(path, ".go") {
			fmt.Println(path)
		}
		return nil
	})
}
