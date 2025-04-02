// Copyright © 2023-2025 Luka Ivanović
// This code is licensed under the terms of the MIT licence (see LICENCE for details)

package diceware

import (
	"math"
	"math/rand"
	"strings"
)

func NewPassphrase() string {
	length := 8
	passphrase := make([]string, length)
	for i := 0; i < length; i += 1 {
		passphrase[i] = words[rollDices()]
	}
	return strings.Join(passphrase, " ")
}

func NewPassphraseWithLength(length int) string {
	if length < 1 || length > 256 {
		length = 8
	}
	passphrase := make([]string, length)
	for i := 0; i < length; i += 1 {
		passphrase[i] = words[rollDices()]
	}
	return strings.Join(passphrase, " ")
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
