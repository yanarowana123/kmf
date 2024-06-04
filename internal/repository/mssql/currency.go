package mssql

import (
	"context"
	"database/sql"
	"github.com/yanarowana123/kmf/internal/model"
	"time"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(connectionString string) Repository {
	db, err := sql.Open("sqlserver", connectionString)
	if err != nil {
		panic(err)
	}
	return Repository{db}
}

func (r Repository) Save(ctx context.Context, currency model.SaveCurrency) error {
	// Пока не придумал какой контекст принять
	con := context.Background()
	_, err := r.db.ExecContext(con, "insert into r_currency (title, code, value, a_date) values (@p1,@p2,@p3, @p4)",
		currency.Title, currency.Code, currency.Value, currency.Date)
	defer r.db.Close()

	return err
}

func (r Repository) List(ctx context.Context, date time.Time, code string) ([]model.GetCurrencyResponse, error) {
	// Пока не придумал какой контекст принять
	con := context.Background()

	var currencies []model.GetCurrencyResponse
	var rows *sql.Rows
	var err error

	if code == "" {
		rows, err = r.db.QueryContext(con, "SELECT title, code, value FROM r_currency WHERE a_date=@p1", date)
	} else {
		rows, err = r.db.QueryContext(con, "SELECT title, code, value FROM r_currency WHERE a_date=@p1 and code = @p2", date, code)
	}

	if err != nil {
		return currencies, err
	}
	defer rows.Close()

	for rows.Next() {
		var currencyResponse model.GetCurrencyResponse
		if err = rows.Scan(&currencyResponse.Title, &currencyResponse.Code, &currencyResponse.Value); err != nil {
			return currencies, err
		}
		currencies = append(currencies, currencyResponse)
	}
	if err = rows.Err(); err != nil {
		return currencies, err
	}
	return currencies, nil
}
