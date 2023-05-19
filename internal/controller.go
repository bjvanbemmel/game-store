package internal

import (
	"encoding/json"
	"net/http"
)

type Data interface {
	struct{} | []struct{} | any | []any
}

type Controller struct{}

func (c Controller) ResponseError(w http.ResponseWriter, err error) {
	var response map[string]string = map[string]string{
		"message": err.Error(),
	}
	raw, _ := json.Marshal(response)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(500)
	w.Write(raw)
}

func (c Controller) Response(w http.ResponseWriter, data Data) {
	raw, err := json.Marshal(data)
	if err != nil {
		c.ResponseError(w, err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(raw)
}
