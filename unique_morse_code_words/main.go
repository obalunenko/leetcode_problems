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
	var uniqueTranslations []string
	for _, word := range words {
		//fmt.Printf("Word: %s\n", word)
		translation := ""
		for _, letter := range word {
			//fmt.Printf(" - letter [%s] -> [ %s ]\n", string(letter), vocabulary[string(letter)])

			translation += vocabulary[string(letter)]

		}
		notSeenBefore := true
		// find notUnique words
		for _, seenTranslation := range uniqueTranslations {
			if translation == seenTranslation {
				notSeenBefore = false
			}
		}
		if notSeenBefore {
			uniqueTranslations = append(uniqueTranslations, translation)
			//fmt.Printf("Unique: %s\n", translated)
		}

	}

	return len(uniqueTranslations)

}
