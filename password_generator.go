package main

import "crypto/rand"

const (
	strLen         int = 32 // Length of random string
	countOfStrings int = 1  // How many string we want to get
)

// Return a new secure random string
func GenerateCryptoSafeString(strLen int) string {
	// Chars collection, as list of ASCII bytes (uint8)
	strPossibleChars := []uint8("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_$")
	lenOfPossibleChars := len(strPossibleChars)

	// Make a buffers in memory
	stringToReturn := make([]byte, strLen)
	randomBytes := make([]byte, strLen)

	iterationsCount := 0

	for {
		// Read a byte
		_, errorOccurred := rand.Read(randomBytes)
		if errorOccurred != nil {
			panic(errorOccurred)
		}

		// For every byte in buffer
		for _, randomByte := range randomBytes {
			// Choose a char
			stringToReturn[iterationsCount] = strPossibleChars[int(randomByte)%lenOfPossibleChars]
			iterationsCount += 1

			// If it will be enough
			if iterationsCount == strLen {

				// Return a random string
				return string(stringToReturn)
			}
		}
	}
}

func main() {
	loopCount := 0

	for loopCount < countOfStrings {
		cryptoSafeString := GenerateCryptoSafeString(strLen)
		println(cryptoSafeString)
		loopCount += 1
	}
}
