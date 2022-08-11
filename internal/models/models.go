package models

import "time"

type RegisterCompay struct {
	Email           string    `json:"email"`
	CompanyPassword string    `json:"company_password"`
	CreatedAt       time.Time `json:"created_at"`
}

type Login struct {
	CompanyUID      string `json:"company_uid"`
	CompanyEmail    string `json:"company_email"`
	CompanyPassword string `json:"company_password"`
}
