package cmd_test

import (
	"testing"

	"github.com/moffgato/ana/cmd"
	"github.com/stretchr/testify/assert"
)

func TestGeneratorAnagrams(t *testing.T) {
	wordSet := map[string]bool{
		"trader": true,
        "retard": true,
		"read":   true,
		"rear":   true,
		"tad":    true,
		"ad":     true,
		"red":    true,
		"ear":    true,
	}

	word := "trader"
	expectedAnagrams := []string{"retard", "trader"}
	anagrams := cmd.GenerateAnagramsFromSubsets(word, wordSet)

	assert.ElementsMatch(t, expectedAnagrams, anagrams, "Anagrams should match")
}

func TestGeneratePermutations(t *testing.T) {
	word := "trader"
	permutations := cmd.GeneratePermutations([]rune(word))

	assert.Contains(t, permutations, "retard", "Permutations should contain 'retard'")
	assert.Contains(t, permutations, "tarred", "Permutations should contain 'tarred'")
}

