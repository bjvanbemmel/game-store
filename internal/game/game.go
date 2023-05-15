package game

import (
	"database/sql"
	"errors"
	"fmt"
	"net/url"

	"github.com/bjvanbemmel/game-store/internal"
	"github.com/bjvanbemmel/game-store/internal/developer"
)

var (
	ErrEmptyField       = errors.New("Field %q may not be empty.")
	ErrFieldOverload    = errors.New("Field %q may not be longer than %d characters.")
	ErrInvalidThumbnail = errors.New("Thumbnail must be an invalid link. %q given.")
)

type Game struct {
	Title       string
	Description string
	Thumbnail   string
	Developer   []*developer.Developer

	internal.Model
}

func Insert(db *sql.DB, game ...*Game) error {
	for _, g := range game {
		if err := validateFields(*g); err != nil {
			return err
		}

		var query string = `
            INSERT INTO games (
                id, title, description, thumbnail
            ) VALUES (
                NULL, ?, ?, ?
            );
        `

		if _, err := db.Exec(query, g.Title, g.Description, g.Thumbnail); err != nil {
			return err
		}
	}

	return nil
}

func First(db *sql.DB, game *Game) error {
	var query string = "SELECT * FROM games LIMIT 1;"

	db.QueryRow(query).Scan(&game.ID, &game.Title, &game.Description, &game.Thumbnail)

	return nil
}

func List(db sql.DB, limit int, games []*Game) error {
	games = nil

	return nil
}

func Paginate(db sql.DB, offset int, limit int, games []*Game) error {
	games = nil

	return nil
}

func validateFields(game Game) error {
	if game.Title == "" {
		return fmt.Errorf(ErrEmptyField.Error(), "Title")
	}

	if len(game.Title) > 255 {
		return fmt.Errorf(ErrFieldOverload.Error(), "Title", 255)
	}

	if game.Description == "" {
		return fmt.Errorf(ErrEmptyField.Error(), "Description")
	}

	if len(game.Description) > 255 {
		return fmt.Errorf(ErrFieldOverload.Error(), "Description", 255)
	}

	if game.Thumbnail == "" {
		return fmt.Errorf(ErrEmptyField.Error(), "Thumbnail")
	}

	if len(game.Thumbnail) > 255 {
		return fmt.Errorf(ErrFieldOverload.Error(), 255)
	}

	if _, err := url.Parse(game.Thumbnail); err != nil {
		return fmt.Errorf(ErrInvalidThumbnail.Error(), game.Thumbnail)
	}

	if len(game.Developer) == 0 {
		return fmt.Errorf(ErrEmptyField.Error(), "Developer")
	}

	return nil
}
