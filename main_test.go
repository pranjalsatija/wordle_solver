package main_test

import (
	"testing"

	wordle "pranj.co/wordle"
)

func BenchmarkEvaluateWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wordle.EvaluateCandidate("mince", "*inc*", "e", "abdfghjklopqrstuvwxyz")
	}
}

func BenchmarkFindCandidates(b *testing.B) {
	for i := 0; i < b.N; i++ {
		scanner, file := wordle.WordScanner()
		wordle.FindCandidates(scanner, "*inc*", "cen", "tarsovph")
		file.Close()
	}
}

func TestEvaluateWord(t *testing.T) {
	if !wordle.EvaluateCandidate("mince", "*inc*", "e", "rtopashv") {
		t.Error("mince failed")
	}

	if !wordle.EvaluateCandidate("wince", "*inc*", "e", "rtopashv") {
		t.Error("wince failed")
	}

	if wordle.EvaluateCandidate("apple", "*****", "", "a") {
		t.Error("apple failed")
	}
}
