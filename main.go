// Copyright (c) 2023 Luka Ivanović
// This code is licensed under MIT licence (see LICENCE for details)

package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) > 1 {
		printUsage()
		os.Exit(1)
	}
	length := 8
	if len(args) == 1 {
		var err error
		length, err = strconv.Atoi(args[0])
		if err != nil {
			printUsage()
			os.Exit(1)
		}
	}
	if length < 1 {
		printUsage()
		os.Exit(1)
	}
	passphrase := make([]string, length)
	for i := 0; i < length; i += 1 {
		passphrase[i] = words[rollDices()]
	}
	fmt.Println(strings.Join(passphrase, " "))
}

func printUsage() {
	fmt.Fprintln(os.Stderr, "Usage: diceware <length>\nDefault length is 8 words, argument <length> is optional and must be positive number.")
}

func rollDice() int {
	return rand.Intn(6) + 1
}

func rollDices() int {
	result := 0
	for i := 0; i < 5; i += 1 {
		result += rollDice() * (int(math.Pow10(5 - i - 1)))
	}
	return result
}
