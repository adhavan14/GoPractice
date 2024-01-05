package utils

import (
	"bufio"
	"os"
)

func loadData(filepath string) ([]string, error) {

	file, err := os.Open(filepath)

	if (err != nil) {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	var data []string
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	return data, nil
}