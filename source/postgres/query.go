package postgres

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Establish a connection with a postgres database.
func ConnectToDatabase() (db *pgxpool.Pool) {
	db, err := pgxpool.New(context.Background(), os.Getenv("PG_URL"))
	if err != nil {
		log.Printf("Error. Could not connect to a database: %v\n", err)
	}
	return db
}

// Set datestyle for the connection.
// This helps avoid errors during insert.
func SetDatestyle(db *pgxpool.Pool) error {
	_, err := db.Exec(context.Background(), "set datestyle = GERMAN, DMY")
	if err != nil {
		log.Printf("Error. Could not set datestyle: %v\n", err)
	}
	return err
}

// Get list of all tables from current database.
func ShowTables(db *pgxpool.Pool) error {
	rows, err := db.Exec(
		context.Background(),
		"SELECT tablename FROM pg_catalog.pg_tables",
	)
	if err != nil {
		log.Printf("Error. Could not load existing tables: %v\n", err)
		return err
	}
	fmt.Println(rows)
	return err
}

// Check if a table with 'tableName' exists in the current database.
func CheckTableExistence(db *pgxpool.Pool, tableName string) (bool, error) {
	query := "SELECT EXISTS (SELECT 1 FROM pg_tables WHERE tablename = '" + tableName + "') AS table_existence;"
	var res bool

	row := db.QueryRow(context.Background(), query)
	err := row.Scan(&res)
	return res, err
}

// Delete a table 'tableName'.
// If it doesn't exist does nothing.
func DeleteTable(db *pgxpool.Pool, tableName string) error {
	query := "DROP TABLE IF EXISTS " + tableName + ";"

	_, err := db.Exec(context.Background(), query)
	if err != nil {
		log.Printf("Error. Could not delete '%v' table: %v\n", tableName, err)
	}
	return err
}

// Renames the table 'oldName' to 'newName'.
func RenameTable(db *pgxpool.Pool, oldName string, newName string) error {
	query := "ALTER TABLE " + oldName + " RENAME TO " + newName + ";"

	_, err := db.Exec(context.Background(), query)
	if err != nil {
		log.Printf("Error. Could not rename '%v' to '%v': %v\n",
			oldName, newName, err)
	}
	return err
}
