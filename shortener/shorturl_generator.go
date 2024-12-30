package shortener

import (
	"crypto/sha256"
	"fmt"
	"github.com/itchyny/base58-go"
	"os"
	"math/big"
)

// sha256Of computes the SHA-256 hash of a given input string and returns the hash as a byte slice.
func sha256Of(input string) []byte {
	// Step 1: Create a new SHA-256 hash algorithm instance
	algorithm := sha256.New()

	// Step 2: Write the input string to the hash algorithm as a byte slice
	// The input string is converted to a byte slice using `[]byte(input)`
	algorithm.Write([]byte(input))

	// Step 3: Compute and return the final hash
	// `Sum(nil)` finalizes the hashing process and returns the hash as a byte slice
	return algorithm.Sum(nil)
}

// base58Encoded encodes a given byte slice into a Base58 string using Bitcoin-style Base58 encoding.
func base58Encoded(bytes []byte) string {
	// `base58.BitcoinEncoding` refers to the standard Base58 encoding used in Bitcoin.
	encoding := base58.BitcoinEncoding

	// `Encode` method converts the byte slice to a Base58-encoded string.
	encoded, err := encoding.Encode(bytes)

	// returns an error and exists if an error is thrown by Encode()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(encoded)
}

// function which generates the actual short link
func GenerateShortLink(initialLink string, userId string) string {
	// computes sha2560 hash of initial link + userID in order to make generated links unique
	urlHashBytes := sha256Of(initialLink + userId)

	// creates a new instance of big.Int because SHA256 hash might be very big
	// then SetBytes() converts the byte slice into an integer
	// Uint64() takes the lower 64 bits of the big number and discards the rest, for simipler processing
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()

	// The `base58Encoded` function encodes the number (converted to a byte slice) into a human-readable Base58 string.
	// The `fmt.Sprintf` formats the number into a decimal string representation before encoding.
	finalString := base58Encoded([]byte(fmt.Sprintf("%d", generatedNumber)))

	// Truncates the Base58-encoded string to 8 characters.
	// This creates a compact and user-friendly short link identifier.
	return finalString[:8]
}
