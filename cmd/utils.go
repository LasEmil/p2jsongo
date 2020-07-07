package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	a "github.com/logrusorgru/aurora"
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

// exists returns whether the given file or directory exists
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}
