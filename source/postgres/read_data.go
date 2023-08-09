package postgres

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/xuri/excelize/v2"
)

func ReadXlsxData(path string) [][]string {
	fd, err := excelize.OpenFile(path)
	if err != nil {
		log.Printf("Error. Could not open .xlsx file: %v\n", err)
		return nil
	}
	rows, err := fd.GetRows(fd.GetSheetName(0))
	if err != nil {
		log.Printf("Error. Could not read .xlsx file: %v\n", err)
		return nil
	}
	return rows[1:]
}

// Read Data from .csv to which tha path points to.
func ReadCsvData(path string) [][]string {
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
