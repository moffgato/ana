package cmd

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

  	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
    "github.com/moffgato/ana/pkg/printer"
)

var (
    words string
    file string
    merge bool
    count int
    dict string
    format string
    output string
    progress bool
)

var anagramCmd = &cobra.Command{
    Use:   "generate",
    Aliases: []string{"g", "gen"},
    Short: "Generate anagrams from word lists",
	Long:  `Generate anagrams from word lists or merged words. You can input words directly or via a file.`,
    Run: AnagramHandler,
}

// read words from dictionary file, returns a proper map for lookups
func readDictionary(filepath string) (map[string]bool, error) {
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
func readWordsFromFile(filepath string) []string {
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
    wordSet, err :=  readDictionary(dict)
    if err != nil {
       fmt.Println(err)
       return
    }

    // read from file or []string args
    if file != "" {
        wordList = readWordsFromFile(file)
    } else if words != "" {
        wordList = strings.Split(words, ",")
    } else {
        fmt.Println("Please provide words or a file to read words from mkay?")
        return
    }

    results := make(map[string][]string)

    var bar *progressbar.ProgressBar
	if progress {
		bar = progressbar.NewOptions(len(wordList),
			progressbar.OptionSetDescription("Generating anagrams..."),
			progressbar.OptionSetTheme(progressbar.Theme{
				Saucer:        "[#]",
				SaucerPadding: " ",
				BarStart:      "[",
				BarEnd:        "]",
			}),
			progressbar.OptionShowCount(),
			progressbar.OptionShowIts(),
			progressbar.OptionSetWidth(20),
			progressbar.OptionSetPredictTime(true),
		)
	}

	for _, word := range wordList {
		word = strings.TrimSpace(word)
		validAnagrams := findValidSubAnagrams(word, wordSet)
		if len(validAnagrams) == 0 {
            results[word] = []string{""}
		} else {
            results[word] = validAnagrams
        }

        if progress {
			bar.Add(1)
		}

	}

    if progress {
        bar.Finish()
        // nukes current line, fixes single word output mess
        fmt.Print("\r\033[K")
    }

    var outputContent string
    switch format {
	case "json":
		printer.JSON(results)
	case "yaml":
		printer.YAML(results)
	case "toml":
		printer.TOML(results)
	case "table":
		printer.Table(results)
	default:
		fmt.Println("Unsupported output format. Supported formats: table, json, yaml, toml")
	}

    if output == "" || output == "stdout" {
		fmt.Print(outputContent)
	} else {
		err := writeToFile(output, outputContent)
		if err != nil {
			fmt.Printf("Error writing to file: %v\n", err)
		} else {
			fmt.Printf("Output written to %s\n", output)
		}
	}

}

// deletes files like rm -rf --no-preserve-root
func writeToFile(filePath, content string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}

func findValidSubAnagrams(word string, wordSet map[string]bool) []string {
	var validAnagrams []string

	// get all possible letter combinations (subsets) of the word
	subsets := generateSubsets([]rune(word))

	// check each subset is a valid word in the dictionary
	for _, subset := range subsets {
		if wordSet[subset] {
			validAnagrams = append(validAnagrams, subset)
		}
	}

	return unique(validAnagrams)
}

// generates all subsets (combinations) of the letters from the given word.
func generateSubsets(chars []rune) []string {
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


func generatePermutations(chars []rune) []string {
    if len(chars) == 1 {
        return []string{string(chars)}
    }

    var permutations []string
    for i := 0; i < len(chars); i++ {
        remaining := append([]rune{}, chars[:i]...)
        remaining = append(remaining, chars[i+1:]...)

        for _, perm := range generatePermutations(remaining) {
            permutations = append(permutations, string(chars[i])+perm)
        }
    }

    return permutations
}

func unique(strings []string) []string {
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
  	anagramCmd.Flags().BoolVarP(&progress, "progress", "p", false, "Display a fancy progress bar")

}

