package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

// returns the tokenized phrase
func tokenize(phrase string) []string {
	return strings.Fields(phrase)
}

// returns only the verbs of a sentence
func verbs(sentence string) []string {
	words := tokenize(sentence)
	verbs := []string{}

	for _, word := range words {
		if strings.HasSuffix(word, "ed") {
			verbs = append(verbs, word)
		}
	}

	return verbs
}

// returns only the nouns of a sentence
func nouns(sentence string) []string {
	words := tokenize(sentence)
	nouns := []string{}

	for _, word := range words {
		if strings.HasPrefix(word, "a") || strings.HasPrefix(word, "an") || strings.HasPrefix(word, "the") {
			nouns = append(nouns, word)
		}
	}

	return nouns
}

type team struct {
	name           string
	points         int
	goalDifference int
}

// returns the leaderboard
func leaderboard(teams []team) []team {
	for i := 0; i < len(teams); i++ {
		for j := i + 1; j < len(teams); j++ {
			if teams[i].points < teams[j].points {
				teams[i], teams[j] = teams[j], teams[i]
			}
		}
	}

	return teams
}
