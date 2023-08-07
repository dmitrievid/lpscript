package postgres

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

const CREATE_RZD_VALUES = ` (
  "Дата отправления"          varchar,
  "Месяц отправления"         varchar,
  "Номер вагона"              varchar,
  "Номер накладной"           varchar,
  "Номер контейнера"          varchar,
  "Префикс контейнера"        varchar,
  "Тип контейнера"            varchar,
  "Вид спецконтейнера"        varchar,
  "Гос отпр"                  varchar,
  "Гос наз"                   varchar,
  "Станц отпр РФ"             varchar,
  "Станц назн РФ"             varchar,
  "Дор отпр РФ"               varchar,
  "Дор назн РФ"               varchar,
  "Регион отпр РФ"            varchar,
  "Регион назн РФ"            varchar,
  "Вид перевозки"             varchar,
  "Груз"                      varchar,
  "Категория отправки"        varchar,
  "Грузоотправитель"          varchar,
  "Грузополучатель"           varchar,
  "Собственник по ЕГРПО"      varchar,
  "Арендатор"                 varchar,
  "Оператор СГТ"              varchar,
  "Плательщик"                varchar,
  "Плательщик пред пор"       varchar,
  "Собственник по вн справ"   varchar,
  "Контейнероотправок"        varchar,
  "Тонн брутто"               varchar,

UNIQUE NULLS NOT DISTINCT (
  "Дата отправления"
  "Месяц отправления"
  "Номер вагона"
  "Номер накладной"
  "Номер контейнера"
  "Префикс контейнера"
  "Тип контейнера"
  "Вид спецконтейнера"
  "Гос отпр"
  "Гос наз"
  "Станц отпр РФ"
  "Станц назн РФ"
  "Дор отпр РФ"
  "Дор назн РФ"
  "Регион отпр РФ"
  "Регион назн РФ"
  "Вид перевозки"
  "Груз"
  "Категория отправки"
  "Грузоотправитель"
  "Грузополучатель"
  "Собственник по ЕГРПО"
  "Арендатор"
  "Оператор СГТ"
  "Плательщик"
  "Плательщик пред пор"
  "Собственник по вн справ"
  "Контейнероотправок"
  "Тонн_брутто"));
`

const INSERT_RZD_VALUES = ` (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6,
  $7,
  $8,
  $9,
  $10,
  $11,
  $12,
  $13,
  $14,
  $15,
  $16,
  $17,
  $18,
  $19,
  $20,
  $21,
  $22,
  $23,
  $24,
  $25,
  $26,
  $27,
  $28,
  $29) ON CONFLICT DO NOTHING;`

func CreateRzdTable(db *pgxpool.Pool, tableName string) error {
	query := "CREATE TABLE IF NOT EXISTS " + tableName + CREATE_RZD_VALUES

	_, err := db.Exec(context.Background(), query)
	if err != nil {
		log.Printf("Error. Could not create '%v' table: %v\n", tableName, err)
	}
	return err
}

func InsertRzdTable(db *pgxpool.Pool, tableName string, rows [][]string) error {
	query := "INSERT INTO " + tableName + INSERT_RZD_VALUES

	for i, row := range rows {
		_, err := db.Exec(context.Background(), query,
			row[0],
			row[1],
			row[2],
			row[3],
			row[4],
			row[5],
			row[6],
			row[7],
			row[8],
			row[9],
			row[10],
			row[11],
			row[12],
			row[13],
			row[14],
			row[15],
			row[16],
			row[17],
			row[18],
			row[19],
			row[20],
			row[21],
			row[22],
			row[23],
			row[24],
			row[25],
			row[26],
			row[27],
			row[28],
		)
		if err != nil {
			log.Printf(
				"Error. Could not insert row number: %v with values: %v to %v: %v\n",
				i, row, tableName, err)
			return err
		}
	}
	return nil
}

const CREATE_STOCKS_VALUES = ` (
"Event Date"          date,
"Owner Name"          varchar(20),
"Terminal"            varchar(50),
"Stock"               varchar(100),
"Planned Arrival"     integer,
"Fact Stock 20"       integer,
"Fact Stock 40"       integer,
"Repair"              integer,
"Reserve Export"      integer,
"Release Cabotage"    varchar(100),
"City Name"           varchar(50),

UNIQUE NULLS NOT DISTINTC (
"Event Date",
"Owner Name",        
"Terminal",
"Stock",
"Planned Arrival",
"Fact Stock 20",
"Fact Stock 40",
"Repair",
"Reserve Export",
"Release Cabotage",
"City Name"));
`

const INSERT_STOCKS_VALUES = ` (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6,
  $7,
  $8,
  $9,
  $10,
  $11) ON CONFLICT DO NOTHING;`

func CreateStocksTable(db *pgxpool.Pool, tableName string) error {
	query := "CREATE TABLE IF NOT EXISTS " + tableName + CREATE_STOCKS_VALUES

	_, err := db.Exec(context.Background(), query)
	if err != nil {
		log.Printf("Error. Could not create '%v' table: %v\n", tableName, err)
	}
	return err
}

func InsertStocksTable(db *pgxpool.Pool, tableName string, rows [][]string) error {
	query := "INSERT INTO " + tableName + INSERT_STOCKS_VALUES

	for i, row := range rows {
		_, err := db.Exec(context.Background(), query,
			row[0],
			row[1],
			row[2],
			row[3],
			row[4],
			row[5],
			row[6],
			row[7],
			row[8],
			row[9],
			row[10],
		)
		if err != nil {
			log.Printf(
				"Error. Could not insert row number: %v with values: %v to %v: %v\n",
				i, row, tableName, err)
			return err
		}
	}
	return nil
}

const CREATE_STAVKI_VALUES = ` (
  "Дата"                date NULL,
	"Месяц"               varchar(8) NULL,
	"Валидность"          date NULL,
	"Компания"            varchar(60) NULL,
	"Тип перевозки"       varchar(10) NULL,
	"Вид перевозки"       varchar(10) NULL,
	"НДС                  smallint NULL,
	"Локация отправления" varchar(50) NULL,
	"Станция отправления" varchar(50) NULL,
	"Локация назначения"  varchar(50) NULL,
	"Станция назначения"  varchar(50) NULL,
	"Принадлежность"      varchar(3) NULL,
	"Футовость"           smallint NULL,
	"Стоимость"           integer NULL,
	"Источник"            varchar(50) NULL,
	"Валюта"              varchar(3) NULL,

UNIQUE NULLS NOT DISTINCT (
  "Дата",
  "Месяц",
  "Валидность",
  "Компания",
  "Тип перевозки",
  "Вид перевозки",
  "НДС,
  "Локация отправления",
  "Станция отправления",
  "Локация назначения",
  "Станция назначения",
  "Принадлежность",
  "Футовость",
  "Стоимость",
  "Источник",
  "Валюта"));
`

const INSERT_STAVKI_VALUES = ` (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6,
  $7,
  $8,
  $9,
  $10,
  $11,
  $12,
  $13,
  $14,
  $15,
  $16) ON CONFLICT DO NOTHING;`

func CreateStavkiTable(db *pgxpool.Pool, tableName string) error {
	query := "CREATE TABLE IF NOT EXISTS " + tableName + CREATE_STAVKI_VALUES

	_, err := db.Exec(context.Background(), query)
	if err != nil {
		log.Printf("Error. Could not create '%v' table: %v\n", tableName, err)
	}
	return err
}

func InsertStavkiTable(db *pgxpool.Pool, tableName string, rows [][]string) error {
	query := "INSERT INTO " + tableName + INSERT_STAVKI_VALUES

	for i, row := range rows {
		_, err := db.Exec(context.Background(), query,
			row[0],
			row[1],
			row[2],
			row[3],
			row[4],
			row[5],
			row[6],
			row[7],
			row[8],
			row[9],
			row[10],
			row[11],
			row[12],
			row[13],
			row[14],
			row[15],
		)
		if err != nil {
			log.Printf(
				"Error. Could not insert row number: %v with values: %v to %v: %v\n",
				i, row, tableName, err)
			return err
		}
	}
	return nil
}
