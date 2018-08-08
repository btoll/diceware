# Diceware

## What is Diceware?

From the [official site][diceware]:

> Diceware™ is a method for picking passphrases that uses dice to select words at random from a special list called the Diceware Word List.

This project is now using the [new word list] from the [EFF], which contain many improvements over the original Diceware word list.

## Installation

`go get github.com/btoll/diceware`

## CLI Usage

    Option | Description
    ------------ | -------------
    -words | The length of the passphrase (defaults to six words).
    -h, --help | Display help.

## Scripting Usage

    Command | Description
    ------------ | -------------
    generate(n) | Generates a passphrase of word length `n` using the Diceware method.

## Ports

- [JavaScript][javascript]

## License

[GPLv3](COPYING)

## Author

Benjamin Toll

[diceware]: http://world.std.com/~reinhold/diceware.html
[new word list]: https://www.eff.org/deeplinks/2016/07/new-wordlists-random-passphrases
[EFF]: https://www.eff.org/
[javascript]: https://github.com/btoll/onf-diceware

