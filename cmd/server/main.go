package main

import (
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{length}/", func(w http.ResponseWriter, r *http.Request) {
		length, err := strconv.Atoi(r.PathValue("length"))
		if err != nil {
			fmt.Fprint(w, "error")
			return
		}
		if length < 0 {
			length = 8
		}
		if length > 255 {
			length = 255
		}
		passphrase := make([]string, length)
		for i := 0; i < length; i += 1 {
			passphrase[i] = words[rollDices()]
		}
		fmt.Fprint(w, strings.Join(passphrase, "-"))
	})
	http.ListenAndServe(":8000", mux)
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
