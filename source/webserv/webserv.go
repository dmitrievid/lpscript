package webserv

import (
	"log"
	ct "logoper_services/postgres/script/controller"
	"net/http"
)

func HandleRequests() {
	http.HandleFunc("/protected/update_rzd", basicAuth(ct.UpdateDatabaseRzd))
	http.HandleFunc("/protected/update_stavki", basicAuth(ct.UpdateDatabaseStavki))
	http.HandleFunc("/protected/update_stocks", basicAuth(ct.UpdateDatabaseStocks))

	log.Fatal(http.ListenAndServe(":444", nil))
}
