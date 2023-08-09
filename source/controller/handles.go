package controller

import (
	"fmt"
	"net/http"
	"os"
	pg "work/logoper_script/postgres"
)

func DeleteCsvFiles(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "Deleting all data files...")
	err := pg.DeleteData(os.Getenv("LP_TEMPDIR"))
	if err != nil {
		fmt.Fprintf(w, "Error. Could not clear the directory: %v\n", err)
		return
	}
	fmt.Fprintln(w, "Successfully deleted all data files.")
}

func UpdateDatabaseStavki(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "Connecting to database...")
	db := pg.ConnectToDatabase()
	if db == nil {
		fmt.Fprintln(w, "Error. Could not connect to the database.")
		return
	}
	defer db.Close()
	fmt.Fprintln(w, "Successfully connected to database.")

	fmt.Fprintln(w, "Setting connection datestyle to DD.MM.YYYY")
	err := pg.SetDatestyle(db)
	if err != nil {
		fmt.Fprintf(w, "Error. Could not set datestyle: %v\n", err)
		return
	}
	fmt.Fprintln(w, "Successfully set connection datestyle.")

	fmt.Fprintln(w, "Checking if the table 'stavki' exists...")
	res, err := pg.CheckTableExistence(db, "stavki")
	if err != nil {
		fmt.Fprintf(w, "Error. Could not check the table's existence: %v\n", err)
		return
	}
	fmt.Fprintf(w, "The table 'stavki' exists? %v\n", res)
	if res != true {
		fmt.Fprintln(w, "Creating table 'stavki'...")
		err := pg.CreateStavkiTable(db, "stavki")
		if err != nil {
			fmt.Fprintf(w, "Error. Could not create table: %v\n", err)
			return
		}
		fmt.Fprintln(w, "Successfully created table.")
	}

	fmt.Fprintln(w, "Reading data from file...")
	rows := pg.ReadXlsxData(os.Getenv("LP_TEMPDIR") + "stavki.xlsx")
	if rows == nil {
		fmt.Fprintln(w, "Error. Could not read data from file.")
		return
	}
	fmt.Fprintln(w, "Successfully read data from file.")

	fmt.Fprintln(w, "Inserting data into the table...")
	err = pg.InsertStavkiTable(db, "stavki", rows)
	if err != nil {
		fmt.Fprintf(w, "Error. Could not insert into the table: %v\n", err)
		return
	}
	fmt.Fprintln(w, "Successfully inserted data into the table.")
}

func UpdateDatabaseStocks(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "Connecting to database...")
	db := pg.ConnectToDatabase()
	if db == nil {
		fmt.Fprintln(w, "Error. Could not connect to the database.")
		return
	}
	defer db.Close()
	fmt.Fprintln(w, "Successfully connected to database.")

	fmt.Fprintln(w, "Setting connection datestyle to DD.MM.YYYY")
	err := pg.SetDatestyle(db)
	if err != nil {
		fmt.Fprintf(w, "Error. Could not set datestyle: %v\n", err)
		return
	}
	fmt.Fprintln(w, "Successfully set connection datestyle.")

	fmt.Fprintln(w, "Checking if the table 'stocks' exists...")
	res, err := pg.CheckTableExistence(db, "stocks")
	if err != nil {
		fmt.Fprintf(w, "Error. Could not check the table's existence: %v\n", err)
		return
	}
	fmt.Fprintf(w, "The table 'stocks' exists? %v\n", res)
	if res != true {
		fmt.Fprintln(w, "Creating table 'stocks'...")
		err := pg.CreateStocksTable(db, "stocks")
		if err != nil {
			fmt.Fprintf(w, "Error. Could not create table: %v\n", err)
			return
		}
		fmt.Fprintln(w, "Successfully created table.")
	}

	fmt.Fprintln(w, "Reading data from file...")
	rows := pg.ReadXlsxData(os.Getenv("LP_TEMPDIR") + "stocks.xlsx")
	if rows == nil {
		fmt.Fprintln(w, "Error. Could not read data from file.")
		return
	}
	fmt.Fprintln(w, "Successfully read data from file.")

	fmt.Fprintln(w, "Inserting data into the table...")
	err = pg.InsertStocksTable(db, "stocks", rows)
	if err != nil {
		fmt.Fprintf(w, "Error. Could not insert into the table: %v\n", err)
		return
	}
	fmt.Fprintln(w, "Successfully inserted data into the table.")
}

func UpdateDatabaseRzd(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "Connecting to database...")
	db := pg.ConnectToDatabase()
	if db == nil {
		fmt.Fprintln(w, "Error. Could not connect to the database.")
		return
	}
	defer db.Close()
	fmt.Fprintln(w, "Successfully connected to database.")

	fmt.Fprintln(w, "Setting connection datestyle to DD.MM.YYYY")
	err := pg.SetDatestyle(db)
	if err != nil {
		fmt.Fprintf(w, "Error. Could not set datestyle: %v\n", err)
		return
	}
	fmt.Fprintln(w, "Successfully set connection datestyle.")

	fmt.Fprintln(w, "Checking if the table 'rzd' exists...")
	res, err := pg.CheckTableExistence(db, "rzd")
	if err != nil {
		fmt.Fprintf(w, "Error. Could not check the table's existence: %v\n", err)
		return
	}
	fmt.Fprintf(w, "The table 'rzd' exists? %v\n", res)
	if res != true {
		fmt.Fprintln(w, "Creating table 'rzd'...")
		err := pg.CreateRzdTable(db, "rzd")
		if err != nil {
			fmt.Fprintf(w, "Error. Could not create table: %v\n", err)
			return
		}
		fmt.Fprintln(w, "Successfully created table.")
	}

	fmt.Fprintln(w, "Reading data from file...")
	rows := pg.ReadXlsxData(os.Getenv("LP_TEMPDIR") + "rzd.xlsx")
	if rows == nil {
		fmt.Fprintln(w, "Error. Could not read data from file.")
		return
	}
	fmt.Fprintln(w, "Successfully read data from file.")

	fmt.Fprintln(w, "Inserting data into the table...")
	err = pg.InsertRzdTable(db, "rzd", rows)
	if err != nil {
		fmt.Fprintf(w, "Error. Could not insert into the table: %v\n", err)
		return
	}
	fmt.Fprintln(w, "Successfully inserted data into the table.")
}
