package cmd

import (
    "fmt"
    "os"
    "github.com/spf13/cobra"
)


var root = &cobra.Command{
    Use: "ana",
    Short: "Anagram CLI to generate anagrams from word list",
  	Long:  `A CLI tool that generates anagrams for words or merged words, with support for input from files.`,
}

func Execute() {
    if err := root.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func init() {

    root.AddCommand(anagramCmd)

}

