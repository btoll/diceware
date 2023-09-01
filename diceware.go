package diceware

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}

func getKey(n uint) string {
	if n < 1 {
		return ""
	}

	// Get ints from CSPRNG.
	num, err := rand.Int(rand.Reader, big.NewInt(6))
	check(err)

	// Need to add 1 so that the random num is between
	// 1 - 6 (inclusive).

	// Convert *big.Int -> string -> integer.
	i, err := strconv.Atoi(num.String())
	check(err)

	// Add 1 to get 1 - 6 inclusive.
	i += 1

	// Convert int back to string.
	key := []string{strconv.Itoa(i)}

	return strings.Join(append(key, getKey(n-1)), "")
}

func getWord() string {
	return Words[getKey(5)]
}

// Generate generates a cryptographically-secure Diceware passphrase.
func Generate(n int, delimiter string) string {
	if n < 1 {
		return ""
	}

	passphrase := []string{getWord()}
	return strings.TrimRight(strings.Join(append(passphrase, Generate(n-1, delimiter)), delimiter), delimiter)
}
