package unique_morse_code_words

type MorseVo1cabulary map[string]string

var morseList = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

var alphabetList string = "abcdefghijklmnopqrstuvwxyz"

var vocabulary = initVocabulary(morseList, alphabetList)

func initVocabulary(morse []string, alphabet string) (result MorseVo1cabulary) {

	return result
}

func uniqueMorseRepresentations(words []string) int {

	return 0

}
