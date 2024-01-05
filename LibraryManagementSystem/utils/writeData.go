package utils

import "os"

func writeData(filepath string, data string) {
	os.WriteFile(filepath, []byte(data), 0644)
}