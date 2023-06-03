package internal

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Db struct {
	*sql.DB
	ConnectionString string
}

func (d *Db) Connect() error {
    var err error

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

    d.DB, err = sql.Open("postgres", d.ConnectionString)
    if err != nil {
        return err
	}

	return nil
}
