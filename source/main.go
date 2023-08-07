package main

import (
	ws "logoper_services/postgres/script/webserv"
)

func main() {
	ws.HandleRequests()
}
