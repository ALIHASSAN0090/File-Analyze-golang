package utils

import (
	"sync"
	"unicode"
)

func GetData(line string, wg *sync.WaitGroup, results chan map[string]int) {
	defer wg.Done()
	counts := map[string]int{"vowels": 0, "capital": 0, "small": 0, "spaces": 0}
	for _, char := range line {
		switch {
		case char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u':
			counts["vowels"]++
			counts["small"]++

		case char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U':
			counts["vowels"]++
			counts["capital"]++

		case unicode.IsUpper(char):
			counts["capital"]++
		case unicode.IsLower(char):
			counts["small"]++
		case unicode.IsSpace(char):
			counts["spaces"]++
		}
	}
	results <- counts
}
