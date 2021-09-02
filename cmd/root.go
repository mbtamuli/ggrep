package cmd

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/spf13/cobra"

	"github.com/mbtamuli/ggrep/grep"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cmd",
	Short: "file pattern searcher",
	Long:  "An implementation of a file pattern searcher similar to grep",
	Run: func(cmd *cobra.Command, args []string) {
		pattern := args[0]
		rootPath := args[1]

		paths, err := grep.ListFiles(rootPath)
		if err != nil {
			log.Fatal(err)
		}

		messages := make(chan string)
		var wg sync.WaitGroup
		for _, path := range paths {
			wg.Add(1)
			go worker(path, pattern, messages, &wg)

		}

		go func() {
			for message := range messages {
				splitString := strings.Split(message, ":")
				path, index, value := splitString[0], splitString[1], splitString[2]
				fmt.Printf("%s:%s> %s\n", path, index, value)
			}

		}()

		wg.Wait()
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func worker(path, pattern string, messages chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	fileContents, _ := grep.ReadLines(path)
	currentFileMatches := grep.Search(fileContents, pattern)
	for index, value := range currentFileMatches {
		messages <- fmt.Sprintf("%s:%d:%s", path, index, value)
	}
}
