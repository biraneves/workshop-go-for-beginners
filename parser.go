package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ParseCsvFile(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error while opening csv file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading records: %w", err)
	}

	return records, nil
}
