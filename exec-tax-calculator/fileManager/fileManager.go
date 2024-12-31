package fileManager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

func ReadTxtFile(fileName string) ([]string, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return nil, errors.New("error opening file")
	}

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		return nil, errors.New("error reading file")
	}

	file.Close()

	return lines, nil
}

func WriteJSONFile(path string, data interface{}) error {
	file, err := os.Create(path)

	if err != nil {
		return errors.New("error creating file")
	}

	err = json.NewEncoder(file).Encode(data)
	if err != nil {
		return errors.New("failed to encode data to Json")
	}

	file.Close()

	return nil
}
