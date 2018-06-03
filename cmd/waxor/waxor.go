package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/xortim/waxor/pkg/waxor"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s string\n", os.Args[0])
		os.Exit(1)
	}

	r, _ := regexp.Compile("^{xor}")
	if r.MatchString(os.Args[1]) {
		fmt.Println(waxor.DecodeXor(os.Args[1]))
	} else {
		fmt.Println(waxor.EncodeXor(os.Args[1]))
	}
}
