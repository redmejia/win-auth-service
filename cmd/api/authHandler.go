package api

import (
	"fmt"
	"net/http"
)

func (a *ApiConfig) AuthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}
