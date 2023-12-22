package fileoperation

import (
	"os"
	"strconv"
)

func ReadFromFile(filename string) (float64, error) {
	byteData, err := os.ReadFile(filename)
	data, err := strconv.ParseFloat(string(byteData), 64)
	return data, err
}