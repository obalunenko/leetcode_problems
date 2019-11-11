package main

type morseVocabulary map[string]string

var morseList = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--",
	"-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

var alphabetList = "abcdefghijklmnopqrstuvwxyz"

func initVocabulary(morse []string, alphabet string) (result morseVocabulary) {
	result = make(morseVocabulary)

	for i, letter := range alphabet {
		result[string(letter)] = morse[i]
	}

	return result
}

func uniqueMorseRepresentations(words []string) int {
	vocabulary := initVocabulary(morseList, alphabetList)

	// make map of translated words
	uniqueTranslations := make(map[string]int)

	for _, word := range words {
		var translation string

		for _, letter := range word {
			translation += vocabulary[string(letter)]
		}

		// find notUnique words
		if _, ok := uniqueTranslations[translation]; !ok {
			uniqueTranslations[translation] = 1
		}
	}

	return len(uniqueTranslations)
}
