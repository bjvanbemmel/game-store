package game

import (
	"github.com/bjvanbemmel/game-store/internal"
	"github.com/bjvanbemmel/game-store/internal/developer"
)

type Game struct {
	Title       string
	Description string
	Thumbnail   string
	Developer   []*developer.Developer

	*internal.Model
}
