package diceware

import (
	"crypto/rand"
	"math/big"
	"strconv"
	"strings"
)

func getKey(n uint) (string, error) {
	if n < 1 {
		return "", nil
	}
	// Get ints from CSPRNG.
	num, err := rand.Int(rand.Reader, big.NewInt(6))
	if err != nil {
		return "", err
	}

	// Need to add 1 so that the random num is between
	// 1 - 6 (inclusive).

	// Convert *big.Int -> string -> integer.
	i, err := strconv.Atoi(num.String())
	if err != nil {
		return "", err
	}

	// Add 1 to get 1 - 6 inclusive.
	i += 1

	// Convert int back to string.
	key := []string{strconv.Itoa(i)}

	word, err := getKey(n - 1)
	if err != nil {
		return "", err
	}
	return strings.Join(append(key, word), ""), nil
}

func getWord() (string, error) {
	word, err := getKey(5)
	if err != nil {
		return "", err
	}
	return Words[word], nil
}

// Generate generates a cryptographically-secure Diceware passphrase.
func Generate(n int, delimiter string) (string, error) {
	if n < 1 {
		return "", nil
	}

	word, err := getWord()
	if err != nil {
		return "", err
	}
	passphrase := []string{word}
	word, err = Generate(n-1, delimiter)
	if err != nil {
		return "", err
	}
	return strings.TrimRight(strings.Join(append(passphrase, word), delimiter), delimiter), nil
}
