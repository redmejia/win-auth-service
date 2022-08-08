package router

import (
	"net/http"
	"win/auth/cmd/api"
)

func Router(a *api.ApiConfig) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/auth", a.AuthHandler)

	return mux
}
