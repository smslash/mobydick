package functions

import (
	"fmt"
	"io/ioutil"
	"os"
)

func DataManipulation(file *os.File) ([][]byte, error) {
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	result := make([]byte, 0)
	for i := 0; i < len(data); i++ {
		if data[i] >= 'A' && data[i] <= 'Z' {
			result = append(result, data[i]+32)
		} else if data[i] >= 'a' && data[i] <= 'z' {
			result = append(result, data[i])
		} else {
			result = append(result, '\n')
		}
	}

	return fields(result), nil
}

func fields(data []byte) [][]byte {
	var result [][]byte
	var start int

	for i := 0; i < len(data); i++ {
		if data[i] == '\n' {
			if i > start {
				field := make([]byte, i-start)
				copy(field, data[start:i])
				result = append(result, field)
			}
			start = i + 1
		}
	}

	if start < len(data) {
		field := make([]byte, len(data)-start)
		copy(field, data[start:])
		result = append(result, field)
	}

	return result
}
