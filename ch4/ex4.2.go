package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func main() {
	var crypto string
	var value string
	var sep string

	for _, arg := range os.Args[1:] {
		if arg == "--sha512" {
			crypto = "sha512"
		} else if arg == "--sha384" {
			crypto = "sha384"
		} else {
			value += sep + arg
			sep = " "
		}
	}

	var result string

	switch crypto {
	case "sha512":
		result = fmt.Sprintf("input: %s\nsha512: %x\n",
			value, sha512.Sum512([]byte(value)))
		break
	case "sha384":
		result = fmt.Sprintf("input: %s\nsha384: %x\n",
			value, sha512.Sum384([]byte(value)))
		break
	default:
		result = fmt.Sprintf("input: %s\nsha256: %x\n",
			value, sha256.Sum256([]byte(value)))
		break
	}

	fmt.Print(result)
}
