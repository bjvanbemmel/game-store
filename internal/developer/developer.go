package developer

import (
	"database/sql"
	"errors"

	"github.com/bjvanbemmel/game-store/internal"
)

var (
	ErrEmptyField    = errors.New("Field %q may not be empty.")
	ErrFieldOverload = errors.New("Field %q may not be longer than %d characters.")
)

type Developer struct {
	Name string

	internal.Model
}

func Insert(db *sql.DB, dev ...*Developer) error {
	var query string = `
        INSERT INTO developers (
            name
        ) VALUES (
            $1
        )
    `

	for _, d := range dev {
		if err := validateFields(*d); err != nil {
			return err
		}

		if _, err := db.Exec(query, d.Name); err != nil {
			return err
		}
	}

	return nil
}

func First(db *sql.DB, dev *Developer) error {
	var query string = "SELECT * FROM developers LIMIT 1;"

	db.QueryRow(query).Scan(&dev.ID, &dev.Name)

	return nil
}

func List(db *sql.DB, limit int, devs *[]*Developer) error {
	var query string = "SELECT * FROM developers LIMIT $1;"

	r, err := db.Query(query, limit)
	if err != nil {
		return err
	}

	for r.Next() {
		var dev Developer = Developer{}
		r.Scan(&dev.ID, &dev.Name)

		*devs = append(*devs, &dev)
	}

	return nil
}

func Paginate(db *sql.DB, page int, limit int, devs *[]*Developer) error {
	if limit == 0 {
		limit = 50
	}

	page -= 1

	var query string = "SELECT * FROM developers LIMIT $1 OFFSET $2;"

	r, err := db.Query(query, limit, page*limit)
	if err != nil {
		return err
	}

	for r.Next() {
		var dev Developer = Developer{}
		r.Scan(&dev.ID, &dev.Name)

		*devs = append(*devs, &dev)
	}

	return nil
}

func validateFields(dev Developer) error {
	if dev.Name == "" {
		return ErrEmptyField
	}

	if len(dev.Name) > 255 {
		return ErrFieldOverload
	}

	return nil
}
