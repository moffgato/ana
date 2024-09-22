package cmd

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/spf13/cobra"
    "github.com/moffgato/ana/pkg/printer"
)

var (
    words string
    file string
    count int
    dict string
    format string
    output string
)

var anagramCmd = &cobra.Command{
    Use:   "generate",
    Aliases: []string{"g", "gen"},
    Short: "Generate anagrams from word lists",
	Long:  `Generate anagrams from word lists. You can input words directly or via a file.`,
    Run: AnagramHandler,
}

// read words from dictionary file, returns a proper map for lookups
func ReadDictionary(filepath string) (map[string]bool, error) {
    file, err := os.Open(filepath)
    if err != nil {
        return nil, fmt.Errorf("error opening dictionary file: %v", err)
    }

    defer file.Close()

    wordSet := make(map[string]bool)
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        word := strings.TrimSpace(scanner.Text())
        wordSet[word] = true
    }

    if err := scanner.Err(); err != nil {
        return nil, fmt.Errorf("error reading dictionary file: %v", err)
    }

    return wordSet, nil

}

// reads wurds from file, returns it as a list of strings
func ReadWordsFromFile(filepath string) []string {
    file, err := os.Open(filepath)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return nil
    }
    defer file.Close()

    var words []string
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        words = append(words, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
        return nil
    }
    return words

}

func AnagramHandler(cmd *cobra.Command, args []string) {

    var wordList []string

    // how can mirrors be real if our eyes aren't real?
    wordSet, err :=  ReadDictionary(dict)
    if err != nil {
       fmt.Println(err)
       return
    }

    // read from file or []string args
    if file != "" {
        wordList = ReadWordsFromFile(file)
    } else if words != "" {
        wordList = strings.Split(words, ",")
    } else {
        fmt.Println("Please provide words or a file to read words from mkay?")
        return
    }

    var results []printer.AnagramOutput

    for _, word := range wordList {
        word = strings.TrimSpace(word)
        validSubsets := FindValidSubAnagrams(word, wordSet)
        anagrams := GenerateAnagramsFromSubsets(word, wordSet)

        results = append(results, printer.AnagramOutput{
            Word:     word,
            Subsets:  validSubsets,
            Anagrams: anagrams,
        })

    }

    // convert results to AnagramResults struct
    finalResults := printer.Output{
        Results: results,
    }

    // output based on selected format
    switch format {
    case "json":
        printer.JSON(finalResults)
    case "yaml":
        printer.YAML(finalResults)
    case "toml":
        printer.TOML(finalResults)
    case "table":
        printer.Table(finalResults)
    default:
        fmt.Println("Unsupported format. Supported formats: table, json, yaml, toml")
        return
    }


    var outputContent string
    if output == "" || output == "stdout" {
		fmt.Print(outputContent)
	} else {
		err := WriteToFile(output, outputContent)
		if err != nil {
			fmt.Printf("Error writing to file: %v\n", err)
		} else {
			fmt.Printf("Output written to %s\n", output)
		}
	}

}

// deletes files á la rm -rf --no-preserve-root
func WriteToFile(filePath, content string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}

func FindValidSubAnagrams(word string, wordSet map[string]bool) []string {
	var validAnagrams []string

	// get all possible letter combinations (subsets) of the word
	subsets := GenerateSubsets([]rune(word))

	// check each subset is a valid word in the dictionary
	for _, subset := range subsets {
		if wordSet[subset] {
			validAnagrams = append(validAnagrams, subset)
		}
	}

	return Unique(validAnagrams)
}

// generates all subsets (combinations) of the letters from the given word.
func GenerateSubsets(chars []rune) []string {
	var subsets []string
	length := len(chars)

	// iterate over all possible combinations of letters
	for i := 1; i < (1 << length); i++ {
		var subset []rune
		for j := 0; j < length; j++ {
			if i&(1<<j) > 0 {
				subset = append(subset, chars[j])
			}
		}
		subsets = append(subsets, string(subset))
	}

	return subsets
}

func GenerateAnagramsFromSubsets(word string, wordSet map[string]bool) []string {
	// get all permutations of the word
	permutations := GeneratePermutations([]rune(word))

	var anagrams []string

	// if the permutation is a valid word or set of valid words
	for _, perm := range permutations {
		if wordSet[perm] {
			anagrams = append(anagrams, perm)
		} else {
			// split into words. check if all are valid
			splitWords := strings.Fields(perm)
			valid := true
			for _, w := range splitWords {
				if !wordSet[w] {
					valid = false
					break
				}
			}
			if valid {
				anagrams = append(anagrams, perm)
			}
		}
	}

	return Unique(anagrams)
}

func GeneratePermutations(chars []rune) []string {

    // when only one char remains, return as single perm
    if len(chars) == 1 {
        return []string{string(chars)}
    }

    var permutations []string
    // iterate chars of input slice
    for i := 0; i < len(chars); i++ {
        // creates slice excluding current char + calls valhalla
        remaining := append([]rune{}, chars[:i]...)
        remaining = append(remaining, chars[i+1:]...)

        // generate permutations for remaining chars
        for _, perm := range GeneratePermutations(remaining) {
            permutations = append(permutations, string(chars[i])+perm)
        }
    }

    return permutations
}


func Unique(strings []string) []string {
    uniqueStrings := make(map[string]bool)
    for _, s := range strings {
        uniqueStrings[s] = true
    }

    var result []string
    for s := range uniqueStrings {
        result = append(result, s)
    }

    sort.Strings(result)
    return result

}

func init() {

  	anagramCmd.Flags().StringVarP(&words, "words", "w", "", "Comma-separated list of words")
	anagramCmd.Flags().StringVarP(&file, "file", "i", "", "File path to read words from")
	anagramCmd.Flags().StringVarP(&dict, "dict", "d", "/usr/share/dict/american-english", "Path to the word list (dictionary)")
	anagramCmd.Flags().StringVarP(&format, "format", "f", "table", "Output format (table, json, yaml, toml)")
	anagramCmd.Flags().StringVarP(&output, "output", "o", "", "Output destination (file path or stdout)")

}

