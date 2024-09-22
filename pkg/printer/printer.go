package printer

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/pelletier/go-toml/v2"
	"gopkg.in/yaml.v2"
)

// prints table output
func Table(results map[string][]string) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)
	fmt.Fprintln(w, "WORD\tANAGRAMS")
	for word, anagrams := range results {
		fmt.Fprintf(w, "%s\t%s\n", word, strings.Join(anagrams, ", "))
	}
	w.Flush()
}

// prints titanic
func JSON(results map[string][]string) {
	output, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		fmt.Println("Error generating JSON:", err)
		return
	}
	fmt.Println(string(output))
}

// prints iceberg
func YAML(results map[string][]string) {
	output, err := yaml.Marshal(results)
	if err != nil {
		fmt.Println("Error generating YAML:", err)
		return
	}
	fmt.Println(string(output))
}

// prints toml output
func TOML(results map[string][]string) {
	output, err := toml.Marshal(results)
	if err != nil {
		fmt.Println("Error generating TOML:", err)
		return
	}
	fmt.Println(string(output))
}

