package api

import (
	"encoding/json"
	"net/http"
	"win/auth/internal/models"
)

func (a *ApiConfig) AuthHandler(w http.ResponseWriter, r *http.Request) {
	var company models.RegisterCompay

	d := json.NewDecoder(r.Body)
	err := d.Decode(&company)
	if err != nil {
		a.ErrorLog.Fatal(err)
	}

	created := a.DB.CreateNewCompany(&company)
	if !created {
		a.ErrorLog.Fatal("unable to create or register company")
	}

}
