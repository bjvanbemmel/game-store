package developer

import (
	"net/http"

	"github.com/bjvanbemmel/game-store/internal"
)

type DeveloperController struct {
	Db *internal.Db

	internal.Controller
}

func (c DeveloperController) Index(w http.ResponseWriter, r *http.Request) {
	var devs []*Developer = []*Developer{}

	if err := c.Db.Ping(); err != nil {
		c.ResponseError(w, err)
	}

	if err := Paginate(c.Db.DB, 1, 10, &devs); err != nil {
		c.ResponseError(w, err)
	}

	c.Response(w, devs)
}
