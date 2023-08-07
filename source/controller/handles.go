package controller

import (
	"fmt"
	pg "logoper_services/postgres/script/postgres"
	"net/http"
)

func UpdateDatabaseStavki(w http.ResponseWriter, _ *http.Request) {
	db := pg.ConnectToDatabase()
	defer db.Close()

	rows := pg.ReadData("../../tmp/stavki.csv")
	pg.InsertStocksTable(db, "stavki", rows)
}

func UpdateDatabaseRzd(w http.ResponseWriter, _ *http.Request) {
	db := pg.ConnectToDatabase()
	defer db.Close()

	rows := pg.ReadData("../../tmp/rzd.csv")
	pg.InsertStocksTable(db, "rzd", rows)
}

func UpdateDatabaseStocks(w http.ResponseWriter, _ *http.Request) {
	db := pg.ConnectToDatabase()
	defer db.Close()

	fmt.Fprintln(w, "Updating database...")
	rows := pg.ReadData("../../tmp/stocks.csv")
	pg.InsertStocksTable(db, "stocks", rows)
	fmt.Fprintln(w, "Database Updated")
}
