package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromValue(filename string) (float64, error) {
	value, err := os.ReadFile(filename)
	if err != nil {
		return 0.0, errors.New("file reading error")
	}

	stringValue := string(value)
	floatValue, err := strconv.ParseFloat(stringValue, 64)
	if err != nil {
		return 0.0, errors.New("converting balance error")
	}

	return floatValue, nil
}

func WriteFloatToFile(value float64, filename string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(filename, []byte(valueText), 0644)
}
