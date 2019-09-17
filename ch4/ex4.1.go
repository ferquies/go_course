package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	hash1 := sha256.Sum256([]byte("x"))
	hash2 := sha256.Sum256([]byte("X"))

	fmt.Printf("%x - %d\n%x - %d\n", hash1, len(hash1), hash2, len(hash2))

	count := 0

	for i := 0; i < len(hash1); i++ {
		if hash1[i] != hash2[i] {
			count++
		}
	}

	fmt.Printf("%d different bytes\n", count)
}
