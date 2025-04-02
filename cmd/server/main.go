// Copyright © 2023-2025 Luka Ivanović
// This code is licensed under the terms of the MIT licence (see LICENCE for details)

package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/luka-hash/diceware/pkg/diceware"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", HandleNewPassphrase)
	mux.HandleFunc("GET /{length}", HandleNewPassphraseWithLength)

	port := os.Getenv("DICEWARE_PORT")
	if port == "" {
		port = "8000"
	}
	fmt.Printf("Listening on localhost:%s\n", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}

func HandleNewPassphrase(w http.ResponseWriter, r *http.Request) {
	passphrase := diceware.NewPassphrase()
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, strings.ReplaceAll(passphrase, " ", "-"))
}

func HandleNewPassphraseWithLength(w http.ResponseWriter, r *http.Request) {
	length, err := strconv.Atoi(r.PathValue("length"))
	if err != nil || length < 1 || length > 256 {
		http.Error(w, "Invalid length", http.StatusBadRequest)
		return
	}
	passphrase := diceware.NewPassphraseWithLength(length)
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, strings.ReplaceAll(passphrase, " ", "-"))
}
