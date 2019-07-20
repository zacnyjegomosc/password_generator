package main

import (
	"sort"
	"strconv"
	"testing"
)

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
		desiredLenOfString             = 16
		maxIterations                  = 100000
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

func TestGenerateCryptoSafeStringCharsDistribution(t *testing.T) {
	const (
		maxTestCount                 = 100   // We want to make this test X times
		checkFirstCharsCount         = 8     // We want to check first X chars from our score list
		digitIsTheBeginningThreshold = 50    // We want to have at least X digits occurs
		desiredLenOfString           = 32    // This is the length of our generated random string
		maxIterations                = 10000 // We want to generate X string for one test
	)

	testCounter := 0
	digitAtTheBeginningOfStringCount := 0
	allTriesCounter := 0

	// For every test iteration...
	for testCounter < maxTestCount {
		i := 0
		// Declare a hash map for chars
		hashMapOfChars := map[string]int{}

		for i < maxIterations {
			// Generate a new random string
			var cryptoSafeString = GenerateCryptoSafeString(desiredLenOfString)

			// Add first char to hash map and increment counter ("char-score")
			hashMapOfChars[string(cryptoSafeString[0])] += 1

			i += 1
		}

		// Sort our char-score list (descending), to build a "winner list"
		numbersMap := map[int][]string{}
		var sliceToSort []int
		for letterKey, letter := range hashMapOfChars {
			numbersMap[letter] = append(numbersMap[letter], letterKey)
		}
		for k := range numbersMap {
			sliceToSort = append(sliceToSort, k)
		}
		sort.Sort(sort.Reverse(sort.IntSlice(sliceToSort)))

		// Check all the "winners"
		charIteration := 1
		for _, k := range sliceToSort {
			allTriesCounter += 1

			// When we check all winners - break the loop and don't check the losers
			if charIteration >= checkFirstCharsCount {
				break
			}

			// When this "winning" char is a digit...
			for _, char := range numbersMap[k] {

				if _, err := strconv.Atoi(char); err == nil {
					// Increment a count of digits in "winners list"
					digitAtTheBeginningOfStringCount += 1
				}
			}

			charIteration += 1

		}
		testCounter += 1
	}

	// If we have too many digits in "winners"...
	if digitAtTheBeginningOfStringCount < digitIsTheBeginningThreshold {
		// This is bad.
		t.Errorf("We have a small chance to get a digit in the beginning. Now it was the only %d for %d tries.",
			digitAtTheBeginningOfStringCount, allTriesCounter)
	}
}

func TestPrepareArgumentsForGenerator(t *testing.T) {
	howManyRandomStringsWeNeed, strLen := PrepareArgumentsForGenerator()
	if howManyRandomStringsWeNeed != defaultCountOfStrings {
		t.Errorf("Invalid desired random strings count.")
	}

	if strLen != defaultStrLen {
		t.Errorf("Invalid desired string length.")
	}
}

func TestRunGenerator(t *testing.T) {
	RunGenerator(defaultCountOfStrings, defaultStrLen)
}
