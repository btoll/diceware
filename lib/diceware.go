/*
Package diceware generates a cryptographically-secure Diceware passphrase
 using the underlying architecture's CSPRNG.

CLI arguments:

    words:
        The length of the passphrase (in words).
*/
package diceware

import (
    "crypto/rand"
    "math/big"
    "strconv"
    "strings"
)

func check(err error) {
    if err != nil {
        panic(err)
    }
}

func getKey(n uint) string {
    if (n < 1) {
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

    return strings.Join(append(key, getKey(n - 1)), "")
}

func getWord() string {
    return Words[getKey(5)]
}

// GetPassphrase generates a cryptographically-secure Diceware passphrase.
func GetPassphrase(n int) string {
    if (n < 1) {
        return ""
    }

    passphrase := []string{getWord()}
    return strings.Join(append(passphrase, GetPassphrase(n - 1)), " ")
}

