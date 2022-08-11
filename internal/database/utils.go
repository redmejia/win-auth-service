package database

import (
	"context"
	"time"
	"win/auth/internal/models"
)

func (d *Postgresql) CreateNewCompay(company *models.RegisterCompay) (created bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	tx, err := d.Db.BeginTx(ctx, nil)
	if err != nil {
		d.ErrorLog.Fatal(err)
	}
	defer tx.Rollback()

	var companyUID string

	row := tx.QueryRowContext(ctx,
		`insert into company_register
			(company_email, company_password, created_at)
		values
			($1, $2, $3) 
		returning company_uid
		`, company.Email, company.CompanyPassword, time.Now())

	err = row.Scan(&companyUID)

	if err != nil {
		d.ErrorLog.Fatal(err)
	}

	_, err = tx.ExecContext(ctx,
		`insert into company_login
			(company_uid, company_email, company_password) 
		values 
			($1, $2, $3)
		`, companyUID, company.Email, company.CompanyPassword,
	)

	err = tx.Commit()
	if err != nil {
		d.ErrorLog.Fatal(err)
	}

	created = true

	return
}
