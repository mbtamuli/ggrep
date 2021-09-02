package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/mbtamuli/ggrep/grep"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cmd",
	Short: "file pattern searcher",
	Long:  "An implementation of a file pattern searcher similar to grep",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Listing files in: %s\n", args[0])
		paths, err := grep.ListFiles(args[0])
		if err != nil {
			log.Fatal(err)
		}

		for _, path := range paths {
			fmt.Printf("%v\n", path)
		}
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cmd.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
