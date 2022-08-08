package main

import (
	"log"
	"net/http"
	"os"
	"win/auth/cmd/api"
	"win/auth/cmd/router"
	"win/auth/internal/database"
)

func main() {
	db, err := database.Connection()
	if err != nil {
		log.Println(err)
	}

	defer db.Close()

	info := log.New(os.Stdout, "INFO\t", log.Ltime|log.Ldate)
	errorl := log.New(os.Stdout, "ERROR\t", log.Ltime|log.Ldate)
	app := api.ApiConfig{
		DB:       database.Postgresql{Db: db, InfoLog: info, ErrorLog: errorl},
		InfoLog:  info,
		ErrorLog: errorl,
	}

	srv := &http.Server{
		Addr:    ":8089",
		Handler: router.Router(&app),
	}

	info.Println("auth service is running at http://localhost:8089/")
	errorl.Fatal(srv.ListenAndServe())
}
