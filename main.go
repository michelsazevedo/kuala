package main

import (
	"log"
	"net/http"

	"github.com/michelsazevedo/kuala/config"
)

func main() {
	app := config.NewApplication()
	log.Fatal(http.ListenAndServe(app.Settings.Server.Port, app.Routes))
}
