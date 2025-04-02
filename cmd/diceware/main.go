// Copyright © 2023-2025 Luka Ivanović
// This code is licensed under the terms of the MIT licence (see LICENCE for details)

package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/luka-hash/diceware/pkg/diceware"
)

func main() {
	args := os.Args[1:]
	if len(args) > 1 {
		printUsage()
		os.Exit(3)
		// "Invalid argument: An incorrect or missing argument was provided to the script."
	}
	length := 8
	if len(args) == 1 {
		var err error
		length, err = strconv.Atoi(args[0])
		if err != nil || length < 1 || length > 256 {
			printUsage()
			os.Exit(3)
			// "Invalid argument: An incorrect or missing argument was provided to the script."
		}
	}
	fmt.Println(diceware.NewPassphraseWithLength(length))
}

func printUsage() {
	fmt.Fprintln(os.Stderr, "Usage: diceware <length>\n\nDefault length is 8 words, argument <length> is optional and must be in range [1, 256].")
}
