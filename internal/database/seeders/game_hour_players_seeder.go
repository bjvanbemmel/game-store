package seeders

import (
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand"

	"github.com/bjvanbemmel/game-store/internal/database"
	"github.com/bjvanbemmel/game-store/internal/dto"
)

type GameHourPlayersSeeder struct{}

const (
	PEAK_HOUR = 20
)

var (
	ErrMinimumInvalid error = errors.New("minimum must be higher than 0")
)

func (s GameHourPlayersSeeder) Seed() error {
	var games []dto.Game

	var query string = `
        INSERT INTO game_hour_players (
            game_id, hour, count
        ) VALUES (
            $1, $2, $3
        );
    `

	rows, err := database.Context.Query(`
        SELECT * FROM games;
    `)
	defer rows.Close()

	if err != nil {
		log.Panic(err)
	}

	for rows.Next() {
		var game dto.Game

		if err := rows.Scan(&game.Id, &game.Title, &game.Price, &game.Thumbnail, &game.Description); err != nil {
			log.Panic(err)
		}

		games = append(games, game)
	}

	for _, g := range games {
		var max int = rand.Intn(50000-5000) + 50000

		for hour := 0; hour < 24; hour++ {
			playerCount, err := s.GeneratePlayerCount(max, hour)
			if err != nil {
				log.Fatal(err)
			}

			if _, err := database.Context.Exec(query, g.Id, hour, playerCount); err != nil {
				return errors.New(fmt.Sprintln(err.Error(), query))
			}
		}
	}

	return nil
}

// Maximum must be higher than 0
func (s GameHourPlayersSeeder) GeneratePlayerCount(maximum int, hour int) (uint, error) {
	var count uint
	var weight float64

	if maximum <= 0 {
		return 0, ErrMinimumInvalid
	}

	weight = math.Abs(1 - (math.Abs(float64(PEAK_HOUR-hour)) / 10))

	count = uint(float64(maximum)*weight) + uint(maximum)/24

	return count, nil
}
