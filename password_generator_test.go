package main

import "testing"

func TestGenerateCryptoSafeStringShortLen(t *testing.T) {
	const desiredLenOfString = 1
	cryptoSafeString := GenerateCryptoSafeString(desiredLenOfString)
	lenOfString := len(cryptoSafeString)
	if lenOfString != desiredLenOfString {
		t.Errorf("Len of string was incorrect (got: %d, want: %d).", lenOfString, desiredLenOfString)
	}
}

func TestGenerateCryptoSafeStringLongLen(t *testing.T) {
	const desiredLenOfString = 256
	cryptoSafeString := GenerateCryptoSafeString(desiredLenOfString)
	lenOfString := len(cryptoSafeString)
	if lenOfString != desiredLenOfString {
		t.Errorf("Len of string was incorrect (got: %d, want: %d).", lenOfString, desiredLenOfString)
	}
}

func TestGenerateCryptoSafeStringDummyCollision(t *testing.T) {
	const (
		desiredLenOfString = 16
		maxIterations = 100000
		defaultValueForStringInHashMap = 0
	)

	i := 0
	// Declare a hash map for random strings
	hashMapOfRandomStrings := map[string]int8{}

	for i < maxIterations {
		// Generate a new random string
		var cryptoSafeString = GenerateCryptoSafeString(desiredLenOfString)

		// Find a random string on hash map
		_, keyWasFound := hashMapOfRandomStrings[cryptoSafeString]

		// If string was found...
		if keyWasFound {

			// This is bad.
			t.Errorf("We have a collision for len=%d (random string is: \"%s\").",
				desiredLenOfString,
				cryptoSafeString)
		}

		// Add string to hash map
		hashMapOfRandomStrings[cryptoSafeString] = defaultValueForStringInHashMap

		i += 1
	}
}