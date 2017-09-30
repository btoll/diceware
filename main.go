package diceware

import (
	"flag"
	"fmt"
)

func main() {
	words := flag.Int("words", 6, "Passphrase length (in words)")
	flag.Parse()

	if *words <= 0 || *words > 500 {
		panic("Value must be between 1 and 500")
	}

	fmt.Println(Generate(*words))
}
