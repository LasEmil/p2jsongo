package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

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
			if response, err := ParseFlat(propertiesFileName, jsonOutputFileName); err != nil {
				panic(err)
			} else {
				fmt.Printf("Wrote %d bytes\n", response)
			}

		},
	}

	var rootCmd = &cobra.Command{Use: "p2json"}
	rootCmd.PersistentFlags().BoolVarP(&flat, "flat", "f", false, "flat parse")
	rootCmd.AddCommand(parse)
	rootCmd.Execute()
}
