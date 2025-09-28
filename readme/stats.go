package readme

import (
	"fmt"
	"slices"
	"sort"
	"strings"
	"unicode"

	a "github.com/jonathon-chew/Aphrodite"
)

var vowels = []rune{'a', 'e', 'i', 'o', 'u', 'y'}

func simpleSyllables(word string) int {
	var result int
	for _, letter := range word {
		if slices.Contains(vowels, letter) {
			result++
		}
	}
	return result
}

func meanLength(StringArray []string) int {
	var wordLength int
	for _, word := range StringArray {
		wordLength += len(word)
	}
	var result int = wordLength / len(StringArray)
	return result
}

func medianLength(StringArray []string) int {
	var result int
	words := StringArray
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	result = len(words) / 2
	return len(words[result])
}

func modeLength(StringArray []string) int {
	var highestCount, result int
	words := StringArray
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	wordLengths := make(map[int]int)
	for _, word := range words {
		wordLengths[len(word)]++
	}
	// a.PrintColour("Green", fmt.Println(wordLengths)
	for key, value := range wordLengths {
		if value >= highestCount {
			result = key
			highestCount = value
		}
	}
	return result
}

func SumArray(intArray []int) int {
	var result int
	for _, number := range intArray {
		result += number
	}
	return result
}

func averageSyllablesPerSentence(content string) int {
	var wordStack []rune
	var sentenceSyllables []int
	endOfSentenceMarkers := []rune{'.', '?', '!'}

	for _, word := range content {
		if !unicode.IsSpace(word) || !slices.Contains(endOfSentenceMarkers, word) { // if the rune is not a space \n \t \r etc
			wordStack = append(wordStack, word)
		} else if unicode.IsSpace(word) { // if the rune is a space figure out the new word!
			sentenceSyllables = append(sentenceSyllables, simpleSyllables(string(wordStack)))
			wordStack = nil
		} else if slices.Contains(endOfSentenceMarkers, word) { // if the rune is a end of sentence marker
			sentenceSyllables = append(sentenceSyllables, simpleSyllables(string(wordStack)))
			wordStack = nil
		} else {
			wordStack, sentenceSyllables = nil, nil
		}
	}

	return SumArray(sentenceSyllables) / len(sentenceSyllables)
}

func NumberOfSentences(Contents string) int {
	var result int
	runes := []rune(Contents)
	for index, sRune := range runes {
		if sRune == '.' || sRune == '?' || sRune == '!' || (sRune == '\n' && !unicode.IsSpace(runes[index-1])) {
			lookForward := index + 1 // Look ahead to the next non-space character
			for lookForward < len(runes) && unicode.IsSpace(runes[lookForward]) {
				lookForward++
			}
			if lookForward >= len(runes) { // Case 1: punctuation is at the very end of text → valid sentence end
				result++
				continue
			}
			if unicode.IsUpper(runes[lookForward]) { // Case 2: next character is uppercase → likely a new sentence
				result++
			}
		}
	}
	return result
}

func wordsOnly(contents string) []string {
	var splitByWord []string = strings.Split(contents, " ")
	for index, word := range splitByWord {
		if word == "\n" {
			splitByWord = slices.Delete(splitByWord, index, index)
		}
	}
	fmt.Print(splitByWord)
	return splitByWord
}

/*
Entry function
*/
func Stats(contents string) error {
	splitByWord := wordsOnly(contents)

	a.PrintColour("Green", fmt.Sprintf("There are %d words in the file\n", len(splitByWord)))
	a.PrintColour("Green", fmt.Sprintf("The origional word count would be %d words in the file\n", len(strings.Split(contents, " "))))
	a.PrintColour("Green", fmt.Sprintf("The mean word size is: %d\n", meanLength(splitByWord)))
	a.PrintColour("Green", fmt.Sprintf("The median word size is: %d\n", medianLength(splitByWord)))
	a.PrintColour("Green", fmt.Sprintf("The mode word size is: %d\n", modeLength(splitByWord)))
	a.PrintColour("Green", fmt.Sprintf("There are %d sentences in the file\n", NumberOfSentences(contents)))
	a.PrintColour("Green", fmt.Sprintf("There are on average %d syllables in each sentence\n", averageSyllablesPerSentence(strings.Join(splitByWord, ""))))

	return nil
}
