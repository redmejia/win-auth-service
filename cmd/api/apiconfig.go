package api

import (
	"log"
	"win/auth/internal/database"
)

type ApiConfig struct {
	DB                database.Postgresql
	InfoLog, ErrorLog *log.Logger
}
