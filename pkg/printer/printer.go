package printer

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

    "github.com/jedib0t/go-pretty/v6/table"
	"github.com/pelletier/go-toml/v2"
	"gopkg.in/yaml.v2"
)

/*

    note to self. this is not great.

*/




type AnagramOutput struct {
    Word string `json:"word" yaml:"word" toml:"word"`
    Subsets []string `json:"subsets" yaml:"subsets" toml:"subsets"`
    Anagrams []string `json:"anagrams" yaml:"anagrams" toml:"anagrams"`
}

type Output struct {
    Results []AnagramOutput `json:"results" yaml:"results" toml:"results"`
}

// prints table output
func Table(results Output) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"WORD", "SUBSETS", "ANAGRAMS"})

	for _, result := range results.Results {
		t.AppendRow(table.Row{
			result.Word,
			strings.Join(result.Subsets, ", "),
			strings.Join(result.Anagrams, ", "),
		})
	}

	t.Render()
}

// prints titanic
func JSON(results Output) {
	output, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		fmt.Println("Error generating JSON:", err)
		return
	}
	fmt.Println(string(output))
}

// prints iceberg
func YAML(results Output) {
	output, err := yaml.Marshal(results)
	if err != nil {
		fmt.Println("Error generating YAML:", err)
		return
	}
	fmt.Println(string(output))
}

// prints toml output
func TOML(results Output) {
	output, err := toml.Marshal(results)
	if err != nil {
		fmt.Println("Error generating TOML:", err)
		return
	}
	fmt.Println(string(output))
}

