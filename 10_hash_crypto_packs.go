// example code with packages hash and crypto
// to hash a string using sha256
package main

import (
	"crypto/sha256"
	"fmt"
	"hash/crc32"
)

func main() {
	// hash a string using sha256
	h := sha256.New()
	h.Write([]byte("test"))
	fmt.Printf("%x\n", h.Sum(nil))

	// hash a string using crc32
	h1 := crc32.NewIEEE()
	h1.Write([]byte("test"))
	v := h1.Sum32()
	fmt.Println(v)
}
