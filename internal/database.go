package internal

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Db struct {
	*sql.DB
	ConnectionString string
}

func (d *Db) Environment(env string) error {
	switch env {
	case "test":
		d.ConnectionString = fmt.Sprintf(`
            host=postgres-test
            user=%s
            password=%s
            dbname=%s
            sslmode=disable`,
			os.Getenv("DB_TEST_USER"),
			os.Getenv("DB_TEST_PASSWORD"),
			os.Getenv("DB_TEST_DATABASE"),
		)
	case "prod":
		d.ConnectionString = fmt.Sprintf(`
            host=postgres-prod
            user=%s
            password=%s
            dbname=%s
            sslmode=disable`,
			os.Getenv("DB_PROD_USER"),
			os.Getenv("DB_PROD_PASSWORD"),
			os.Getenv("DB_PROD_DATABASE"),
		)
	default:
		return errors.New("Given environment does not exist.")
	}

	return nil
}

func (d *Db) Connect() error {
	var err error
	d.DB, err = sql.Open("postgres", d.ConnectionString)
	if err != nil {
		return err
	}

	return nil
}
