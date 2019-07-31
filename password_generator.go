package main

import (
	"crypto/rand"
	"flag"
)

const (
	defaultStrLen         int    = 32 // Length of random string
	defaultCountOfStrings int    = 1  // How many string we want to get
	possibleChars         string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_$"
)

// GenerateCryptoSafeString returns a new secure random string
func GenerateCryptoSafeString(strLen int) string {
	// Chars collection, as list of ASCII bytes (uint8)
	strPossibleChars := []uint8(possibleChars)
	lenOfPossibleChars := len(strPossibleChars)

	// Make a buffers in memory
	stringToReturn := make([]byte, strLen)
	randomBytes := make([]byte, strLen)

	// Read a random bytes
	_, errorOccurred := rand.Read(randomBytes)
	if errorOccurred != nil {
		panic(errorOccurred)
	}

	charNumber := 0

	// For every byte in buffer
	for _, randomByte := range randomBytes {

		// Choose a char and add it to stringToReturn
		stringToReturn[charNumber] = strPossibleChars[int(randomByte)%lenOfPossibleChars]
		charNumber += 1
	}

	// Return a random string
	return string(stringToReturn)
}

// PrepareArgumentsForGenerator reads command-line flags or set defaults values
func PrepareArgumentsForGenerator() (int, int) {
	// Read command-line flags or set defaults values
	howManyRandomStringsWeNeed := flag.Int("count", defaultCountOfStrings, "How many strings do you need?")
	strLen := flag.Int("len", defaultStrLen, "What length of string do you need?")

	flag.Parse()

	return *howManyRandomStringsWeNeed, *strLen
}

// RunGenerator evaluates all high-level logic
func RunGenerator(howManyRandomStringsWeNeed int, strLen int) {
	loopCount := 0

	for loopCount < howManyRandomStringsWeNeed {
		cryptoSafeString := GenerateCryptoSafeString(strLen)
		println(cryptoSafeString)
		loopCount += 1
	}
}

func main() {
	howManyRandomStringsWeNeed, strLen := PrepareArgumentsForGenerator()
	RunGenerator(howManyRandomStringsWeNeed, strLen)
}
