package fileoperation

import (
	"fmt"
	"os"
)

func WriteToFile(filename string, data float64) {
	os.WriteFile(filename, []byte(fmt.Sprint(data)), 0644)
}