package file_io

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(fileName string, value float64) {
	floatText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(floatText), 0644)
}

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 0.0, errors.New("failed to read balance file")
	}

	floatText := string(data)
	value, err := strconv.ParseFloat(floatText, 64)
	if err != nil {
		return 0.0, errors.New("failed to parse balance file")
	}

	return value, nil
}
