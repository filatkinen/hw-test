package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type wordsCount struct {
	word  string
	count int
}

const TaskWithAsteriskIsCompleted = false

// Transform to lowercase for Latin and Russian alphabets using pointer to avoid excessive coping.
func toLower(ch *rune) {
	if *ch >= 'А' && *ch <= 'Я' || *ch >= 'A' && *ch <= 'Z' {
		*ch += 32
	}
}

// Checking if char belongs to  Latin or Russian alphabets and '-'.
func isCharInAlphabetAndDash(ch *rune) bool {
	return *ch >= 'а' && *ch <= 'я' || *ch == '-' ||
		*ch >= 'А' && *ch <= 'Я' || *ch >= 'a' && *ch <= 'z' ||
		*ch >= 'A' && *ch <= 'Z'
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
		thisCharInAlphabetAndDash := isCharInAlphabetAndDash(&txt[i])
		toLower(&txt[i])
		switch {
		case (i == len(txt)-1) && thisCharInAlphabetAndDash:
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

func Top10(text string) []string {
	if len(text) == 0 {
		return nil
	}

	var wordsMap map[string]int
	switch TaskWithAsteriskIsCompleted {
	case false:
		wordsMap = fromTextToWordsMap(text)
	case true:
		wordsMap = fromTextToWordsMapWithAsterisk(text)
	}

	// Coping Map values to slice with structure wordCount{word,count}
	wordsCountList := make([]wordsCount, 0, len(wordsMap))
	for word, count := range wordsMap {
		wordsCountList = append(wordsCountList, wordsCount{
			word:  word,
			count: count,
		})
	}

	// 1- sorting slice by frequency
	sort.Slice(wordsCountList, func(i, j int) bool {
		return wordsCountList[i].count > wordsCountList[j].count
	})

	// 2-sorting slice by lexicographically
	idx := 0
	tmpSlice := []wordsCount(nil)
	if len(wordsCountList) > 0 {
		for i := 1; i < len(wordsCountList); i++ {
			needSort := false
			if wordsCountList[idx].count != wordsCountList[i].count {
				tmpSlice = wordsCountList[idx:i]
				needSort = true
			}
			if (i == len(wordsCountList)-1) && wordsCountList[idx].count == wordsCountList[i].count {
				tmpSlice = wordsCountList[idx : i+1]
				needSort = true
			}
			if needSort {
				sort.Slice(tmpSlice, func(i, j int) bool {
					return tmpSlice[i].word < tmpSlice[j].word
				})
				idx = i
			}
		}
	}

	// coping first 10 values to the result slice checking if source's size  has less than 10
	l := len(wordsCountList)
	if l > 10 {
		l = 10
	}
	wordsList := make([]string, 0, l)
	for i := 0; i < l; i++ {
		wordsList = append(wordsList, wordsCountList[i].word)
	}

	return wordsList
}
