package webserv

import (
	"log"
	"net/http"
	ct "work/logoper_script/controller"
)

func HandleRequests() {
	http.HandleFunc("/protected/update_rzd", basicAuth(ct.UpdateDatabaseRzd))
	http.HandleFunc("/protected/update_stavki", basicAuth(ct.UpdateDatabaseStavki))
	http.HandleFunc("/protected/update_stocks", basicAuth(ct.UpdateDatabaseStocks))
	http.HandleFunc("/protected/delete_files", basicAuth(ct.DeleteCsvFiles))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
