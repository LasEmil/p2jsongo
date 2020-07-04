package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

var ErrMalformedFile = errors.New("File is malformed")
var ErrCreatingFIle = errors.New("Error when creating file")

// ParseFlat parses properties file to json only on one level, not going deep into the properties
func ParseFlat(pFileName, jsonFileName string) (int, error) {
	m := make(map[string]string)
	file, err := os.Open(pFileName)
	skipLineCounter := 0
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		eqIdx := strings.Index(line, "=")
		if eqIdx > -1 {
			key := line[:eqIdx]
			value := line[eqIdx+1:]
			m[key] = value
		} else {
			skipLineCounter = skipLineCounter + 1
		}
	}
	json, err := jsoniter.Marshal(&m)
	if err != nil {
		return 0, err
	}

	newFile, err := os.Create(jsonFileName)
	if err != nil {
		return 0, ErrCreatingFIle
	}
	defer newFile.Close()
	w := bufio.NewWriter(newFile)
	writeBuffer, err := w.WriteString(string(json))
	if err != nil {
		return 0, err
	}
	w.Flush()
	if skipLineCounter > 0 {
		fmt.Printf("Skipped %d lines.\n", skipLineCounter)
	}
	return writeBuffer, nil
}

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

// MyMap is empty map interface
type MyMap map[string]interface{}

func createPath(m MyMap, path string, value string) MyMap {
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

// ParseDeep function parses the java's properties file into nested json
func ParseDeep(pFileName, jsonFileName string) (int, error) {
	m := make(MyMap)
	file, err := os.Open(pFileName)
	skipLineCounter := 0
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		eqIdx := strings.Index(line, "=")
		if eqIdx > -1 {
			key := line[:eqIdx]
			value := line[eqIdx+1:]
			createPath(m, key, value)
		} else {
			skipLineCounter = skipLineCounter + 1
		}
	}
	json, err := jsoniter.Marshal(&m)
	if err != nil {
		return 0, err
	}

	newFile, err := os.Create(jsonFileName)
	if err != nil {
		return 0, ErrCreatingFIle
	}
	defer newFile.Close()
	w := bufio.NewWriter(newFile)
	writeBuffer, err := w.WriteString(string(json))
	if err != nil {
		return 0, err
	}
	w.Flush()
	if skipLineCounter > 0 {
		fmt.Printf("Skipped %d lines.\n", skipLineCounter)
	}
	return writeBuffer, nil
}
