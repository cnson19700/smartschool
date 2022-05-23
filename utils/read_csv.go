package utils

import (
	"encoding/csv"
	"io"
	"os"
)

func GetColumnDataCSV(filePath string, column int) []string {
	f, err := os.Open(filePath)
	if err != nil {
		return nil
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	data := make([]string, 0)
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return data
		}
		data = append(data, record[column])
	}
	return data
}
