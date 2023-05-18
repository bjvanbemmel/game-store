package developer

import (
	"encoding/json"
	"net/http"

	"github.com/bjvanbemmel/game-store/internal"
)

type DeveloperController struct {
	Db *internal.Db
}

func (c DeveloperController) Index(w http.ResponseWriter, r *http.Request) {
	var devs []*Developer = []*Developer{}

	w.Header().Add("Content-Type", "application/json")

	if err := c.Db.Ping(); err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}

	if err := Paginate(c.Db.DB, 1, 10, &devs); err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}

	raw, err := json.Marshal(devs)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}

	w.Write(raw)
}
