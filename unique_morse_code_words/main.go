package unique_morse_code_words

type MorseVocabulary map[string]string

var morseList = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

var alphabetList = "abcdefghijklmnopqrstuvwxyz"

func initVocabulary(morse []string, alphabet string) (result MorseVocabulary) {
	result = make(MorseVocabulary)

	for i, letter := range alphabet {
		result[string(letter)] = morse[i]
		//fmt.Printf("%d: %s -> %s\n", i, string(letter), morse[i])
	}

	return result
}

func uniqueMorseRepresentations(words []string) int {
	vocabulary := make(MorseVocabulary)
	vocabulary = initVocabulary(morseList, alphabetList)

	// make map of translated words
	uniqueTranslations := make(map[string]int)
	for _, word := range words {

		var translation string

		for _, letter := range word {

			translation += vocabulary[string(letter)]

		}
		// find notUnique words
		_, ok := uniqueTranslations[translation]
		if !ok {
			uniqueTranslations[translation] = 1
		}

	}

	return len(uniqueTranslations)

}
