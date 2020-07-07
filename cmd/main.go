package cmd

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	a "github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

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

// Execute function executes the program
func Execute() {
	flat := false
	var parse = &cobra.Command{
		Use:   "parse [properties file or directory to parse] [destination filename (optional)]",
		Short: "Parse properties file to json",
		Long:  `Parse java's properties file format to json. You can parse single file or entire directory. When parsing single file you have an option to add second argument: output filename, when parsing entire directory this argument is skipped and all files will be named like the source files`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			source := args[0]
			sourceExists, err := exists(source)
			if err != nil {
				panic(err)
			}
			if sourceExists {
				defer elapsed("Parsing")()

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
			} else {
				fmt.Printf("%s: No such file or directory\n", source)
			}
		},
	}

	var rootCmd = &cobra.Command{Use: "p2jsongo"}
	rootCmd.PersistentFlags().BoolVarP(&flat, "flat", "f", false, "flat parse")
	rootCmd.AddCommand(parse)
	rootCmd.Execute()
}
