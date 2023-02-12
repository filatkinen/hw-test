package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

// Transform to lowercase for Latin and Russian alphabets using pointer to avoid excessive coping.
func toLower(ch rune) rune {
	if ch >= 'А' && ch <= 'Я' || ch >= 'A' && ch <= 'Z' {
		ch += 32
	}
	return ch
}

// Checking if char belongs to  Latin or Russian alphabets and '-'.
func isCharInAlphabetAndDash(ch rune) bool {
	return ch >= 'а' && ch <= 'я' || ch == '-' ||
		ch >= 'А' && ch <= 'Я' || ch >= 'a' && ch <= 'z' ||
		ch >= 'A' && ch <= 'Z'
}

func fromTextToWordsMap(text string) map[string]int {
	// word's average is about 10 chars
	wordsMap := make(map[string]int, len(text)/10)
	words := strings.Fields(text)
	for _, word := range words {
		wordsMap[word]++
	}
	return wordsMap
}

func fromTextToWordsMapWithAsterisk(text string) map[string]int {
	// word's average is about 10 chars
	wordsMap := make(map[string]int, len(text)/10)
	wasAlphabet := false
	needToAdd := false
	prevIdx := 0
	thisIdx := 0
	txt := []rune(text)
	for i := 0; i < len(txt); i++ {
		thisCharInAlphabetAndDash := isCharInAlphabetAndDash(txt[i])
		txt[i] = toLower(txt[i])
		switch {
		case (i == len(txt)-1) && thisCharInAlphabetAndDash:
			if !wasAlphabet {
				prevIdx = i
			}
			needToAdd = true
			thisIdx = i + 1
		case thisCharInAlphabetAndDash:
			if !wasAlphabet {
				prevIdx = i
				wasAlphabet = true
			}
		case !thisCharInAlphabetAndDash:
			if wasAlphabet {
				needToAdd = true
				wasAlphabet = false
				thisIdx = i
			}
		}
		if needToAdd {
			wordsMap[string(txt[prevIdx:thisIdx])]++
			needToAdd = false
		}
	}
	// delete -
	delete(wordsMap, "-")
	return wordsMap
}

func Top10(text string, withAsterisk bool) []string {
	if len(text) == 0 {
		return nil
	}

	var wordsMap map[string]int
	if withAsterisk {
		wordsMap = fromTextToWordsMapWithAsterisk(text)
	} else {
		wordsMap = fromTextToWordsMap(text)
	}

	// making slice to be able sorting
	wordsSlice := make([]string, 0, len(wordsMap))
	for word := range wordsMap {
		wordsSlice = append(wordsSlice, word)
	}

	// 1- sorting slice by frequency and then lexicographically
	sort.Slice(wordsSlice, func(i, j int) bool {
		if wordsMap[wordsSlice[i]] == wordsMap[wordsSlice[j]] {
			return wordsSlice[i] < wordsSlice[j]
		}
		return wordsMap[wordsSlice[i]] > wordsMap[wordsSlice[j]]
	})

	// return first 10 values to the result slice checking if source's size  has less than 10
	l := len(wordsSlice)
	if l > 10 {
		l = 10
	}
	return wordsSlice[:l]
}
