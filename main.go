package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

func EvaluateCandidate(word string, hint string, unplacedCharacters string, rejectedCharacters string) bool {
	if utf8.RuneCountInString(word) != 5 {
		return false
	}

	for i := 0; i < 5; i++ {
		if hint[i] != '*' && word[i] != hint[i] {
			return false
		}

		if strings.Contains(rejectedCharacters, string(word[i])) {
			return false
		}

		for _, unplacedCharacter := range unplacedCharacters {
			if !strings.ContainsRune(word, unplacedCharacter) {
				return false
			}
		}
	}

	return true
}

func FindCandidates(wordScanner *bufio.Scanner, hint string, unplacedCharacters string, rejectedCharacters string) []string {
	potentialWords := make([]string, 0)

	for wordScanner.Scan() {
		word := strings.TrimSpace(strings.ToLower(wordScanner.Text()))

		if EvaluateCandidate(word, hint, unplacedCharacters, rejectedCharacters) {
			potentialWords = append(potentialWords, word)
		}
	}

	return potentialWords
}

func WordScanner() (*bufio.Scanner, *os.File) {
	file, err := os.Open("/usr/share/dict/words")

	if err != nil {
		panic(err)
	}

	return bufio.NewScanner(file), file
}

func main() {
	scanner, file := WordScanner()
	hint := flag.String("hint", "*****", "A string containing the known characters in the word. Unknown characters should be represented as asterisks.")
	unplacedCharacters := flag.String("unplaced", "", "A string containing the characters that are known to be in the word, but aren't placed.")
	rejectedCharacters := flag.String("rejected", "", "A string containing the characters that have been rejected from the word.")

	flag.Parse()

	if utf8.RuneCountInString(*hint) != 5 {
		log.Fatalf("The hint must contain exactly 5 characters.")
	}

	for _, candidate := range FindCandidates(scanner, *hint, *unplacedCharacters, *rejectedCharacters) {
		fmt.Println(candidate)
	}

	file.Close()
}
