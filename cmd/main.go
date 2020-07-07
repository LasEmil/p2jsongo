package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}

// GetFileName returns file name without extension
func getFileName(arg string) string {
	dotIdx := strings.Index(arg, ".")
	fileName := arg[:dotIdx]
	return fileName
}
func parseSingleFile(args []string, source string, outputFileName string, flat bool) {
	defer elapsed("Parsing")()

	if len(args) == 1 {
		fileName := getFileName(source)
		outputFileName = fmt.Sprintf("%s.%s", fileName, "json")
	} else {
		outputFileName = args[1]
	}

	if response, err := Parse(source, outputFileName, flat); err != nil {
		panic(err)
	} else {
		fmt.Printf("Wrote %d bytes\n", response)
	}
}

// Execute function executes the program
func Execute() {
	flat := false
	var parse = &cobra.Command{
		Use:   "parse [properties file to parse] [destination file]",
		Short: "Parse properties file to json",
		Long:  `Parse java's properties file format to json`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			source := args[0]
			var outputFileName string
			fi, err := os.Stat(source)
			if err != nil {
				fmt.Println(err)
				return
			}

			switch mode := fi.Mode(); {
			case mode.IsDir():
				// Entire directory parse
				files, err := ioutil.ReadDir(source)
				if err != nil {
					log.Fatal(err)
				}
				for _, f := range files {
					extension := filepath.Ext(f.Name())
					if extension == ".properties" {
						fileName := getFileName(f.Name())
						path := fmt.Sprintf("%s/%s", source, f.Name())
						outputPath := fmt.Sprintf("%s/%s.%s", source, fileName, "json")
						fmt.Println(path)
						fmt.Println(outputPath)
					} else {
						fmt.Printf("Not the properties file: %s\n", f.Name())
					}
				}
			case mode.IsRegular():
				// Single file parse
				parseSingleFile(args, source, outputFileName, flat)
			}
		},
	}

	var rootCmd = &cobra.Command{Use: "p2jsongo"}
	rootCmd.PersistentFlags().BoolVarP(&flat, "flat", "f", false, "flat parse")
	rootCmd.AddCommand(parse)
	rootCmd.Execute()
}
