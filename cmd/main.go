package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	a "github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		timeSinceStart := time.Since(start)
		if timeSinceStart > time.Duration(time.Microsecond*500000) {
			fmt.Printf("%s took %v\n", what, a.Yellow(timeSinceStart))
		} else {
			fmt.Printf("%s took %v\n", what, a.Green(timeSinceStart))

		}
	}
}

// GetFileName returns file name without extension
func getFileName(arg string) string {
	dotIdx := strings.Index(arg, ".")
	fileName := arg[:dotIdx]
	return fileName
}
func parseSingleFile(source string, outputFileName string, flat bool) (int, error) {
	response, err := Parse(source, outputFileName, flat)
	if err != nil {
		return 0, err
	}
	return response, nil

}
func parseDirectory(source string, flat bool) (int, error) {
	files, err := ioutil.ReadDir(source)
	bytesWritten := 0
	skippedFiles := 0
	if err != nil {
		return 0, err
	}
	if len(files) == 0 {
		return 0, fmt.Errorf("There is no files in the directory")
	}
	for _, f := range files {
		extension := filepath.Ext(f.Name())
		if extension == ".properties" {
			fileName := getFileName(f.Name())
			path := fmt.Sprintf("%s/%s", source, f.Name())
			outputPath := fmt.Sprintf("%s/%s.%s", source, fileName, "json")

			fileBytesWritten, err := parseSingleFile(path, outputPath, flat)
			if err != nil {
				return 0, err
			}
			bytesWritten = bytesWritten + fileBytesWritten
		} else {
			skippedFiles++
		}
	}
	if skippedFiles > 0 {
		fmt.Printf("Skipped %d files in directory %s (not \".properties\" file)\n", a.Red(skippedFiles), source)
	}
	return bytesWritten, nil
}

func isDir(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
		return false
	}
	switch mode := fi.Mode(); {
	case mode.IsDir():
		// Entire directory parse
		return true

	case mode.IsRegular():
		// Single file parse
		return false
	}
	return false
}

// Execute function executes the program
func Execute() {
	flat := false
	var parse = &cobra.Command{
		Use:   "parse [properties file or directory to parse] [destination filename (optional)]",
		Short: "Parse properties file to json",
		Long:  `Parse java's properties file format to json. You can parse single file or entire directory. When parsing single file you have an option to add second argument: output filename, when parsing entire directory this argument is skipped and all files will be named like the source files`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			defer elapsed("Parsing")()
			source := args[0]

			if isDir(source) {
				bytesWritten, err := parseDirectory(source, flat)
				if err != nil {
					panic(err)
				}
				fmt.Printf("Bytes written %d\n", a.Cyan(bytesWritten))
			} else {
				var outputFileName string
				if len(args) == 1 {
					fileName := getFileName(source)
					outputFileName = fmt.Sprintf("%s.%s", fileName, "json")
				} else {
					outputFileName = args[1]
				}
				bytesWritten, err := parseSingleFile(source, outputFileName, flat)
				if err != nil {
					panic(err)
				}
				fmt.Printf("Written %d bytes.\n", a.Cyan(bytesWritten))
			}
		},
	}

	var rootCmd = &cobra.Command{Use: "p2jsongo"}
	rootCmd.PersistentFlags().BoolVarP(&flat, "flat", "f", false, "flat parse")
	rootCmd.AddCommand(parse)
	rootCmd.Execute()
}
