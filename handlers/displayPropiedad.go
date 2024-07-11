package handlers

import (
	"net/http"
	"strconv"

	"PagProp/db"
	"PagProp/models"
	"PagProp/views/displayPropiedad"
	"github.com/go-chi/chi/v5"
)

func HandleDisplayPropiedad(w http.ResponseWriter, r *http.Request) error {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}

	database := db.Create()
	propiedad := &models.Propiedad{}

	err = database.NewSelect().Model(propiedad).Where("id = ?", id).Scan(r.Context())
	if err != nil {
		return err
	}

	return Render(w, r, displayPropiedad.DisplayPropiedad(propiedad))
}
