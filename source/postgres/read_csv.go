package postgres

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

// Read Data from .csv to which tha path points to.
func ReadData(path string) [][]string {
	fd, err := os.Open(path)
	if err != nil {
		log.Printf("Error. Could not load .csv file: %v\n", err)
		return nil
	}
	defer fd.Close()

	// Instantiate a new reader and set the delimiter
	reader := csv.NewReader(fd)
	reader.Comma = ';'

	// Skip first line with column names
	_, err = reader.Read()
	if err != nil {
		log.Printf("Error. Could not read .csv file: %v\n", err)
		return nil
	}

	rows, err := reader.ReadAll()
	if err != nil {
		log.Printf("Error. Could not read .csv file: %v\n", err)
		return nil
	}
	return rows
}

func DeleteData(dirPath string) error {
	list, err := os.ReadDir(dirPath)
	if err != nil {
		log.Printf("Error. Could not read files from tmp directory: %v\n", err)
		return err
	}
	for i, val := range list {
		fmt.Println(i, val)
	}
	return nil
}
