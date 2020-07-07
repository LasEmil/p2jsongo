package cmd

import (
	"bufio"
	"errors"
	"os"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

var errMalformedFile = errors.New("File is malformed")
var errCreatingFile = errors.New("Error while creating file")
var errNoSuchFile = errors.New("Properties file: No such file or directory")

// MyMap is empty map interface
type MyMap map[string]interface{}

func createPath(m map[string]interface{}, path string, value string) MyMap {
	pathSlice := strings.Split(path, ".")
	current := m
	for len(pathSlice) > 1 {
		var head = pathSlice[0]
		var tail = pathSlice[1:]
		pathSlice = tail
		if _, ok := current[head]; !ok {
			current[head] = make(MyMap)
		}
		if typeof(current[head]) != "string" {
			current = current[head].(MyMap)
		}
	}
	current[pathSlice[0]] = value
	return m
}

// Parse function parses the java's properties file into nested json
func Parse(input, output string, flat bool) (int, error) {
	m := make(map[string]interface{})
	file, err := os.Open(input)
	if err != nil {
		return 0, errNoSuchFile
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			// Removing comments
			if string(line[0]) != "#" {
				eqIdx := strings.Index(line, "=")
				if eqIdx > -1 {
					key := line[:eqIdx]
					value := line[eqIdx+1:]
					if flat {
						m[key] = value
					} else {
						createPath(m, key, value)
					}
				}
			}
		}

	}
	json, err := jsoniter.Marshal(&m)
	if err != nil {
		return 0, err
	}

	newFile, err := os.Create(output)
	if err != nil {
		return 0, errCreatingFile
	}
	defer newFile.Close()
	w := bufio.NewWriter(newFile)
	writeBuffer, err := w.WriteString(string(json))
	if err != nil {
		return 0, err
	}
	w.Flush()

	return writeBuffer, nil
}
