package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}

// Execute function executes the program
func Execute() {
	flat := false
	var parse = &cobra.Command{
		Use:   "parse [properties file to parse] [destination file]",
		Short: "Parse properties file to json",
		Long:  `Parse java's properties file format to json`,
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			propertiesFileName := args[0]
			jsonOutputFileName := args[1]
			defer elapsed("Parsing")()

			if response, err := Parse(propertiesFileName, jsonOutputFileName, flat); err != nil {
				panic(err)
			} else {
				fmt.Printf("Wrote %d bytes\n", response)
			}

		},
	}

	var rootCmd = &cobra.Command{Use: "p2jsongo"}
	rootCmd.PersistentFlags().BoolVarP(&flat, "flat", "f", false, "flat parse")
	rootCmd.AddCommand(parse)
	rootCmd.Execute()
}
