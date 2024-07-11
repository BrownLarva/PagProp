package handlers

import (
	"context"
	"net/http"

	"PagProp/db"
	"PagProp/models"
	"PagProp/views/home"

	"github.com/uptrace/bun"
)

func fetchPropiedades(db *bun.DB) []models.Propiedad {
	var propiedades []models.Propiedad
	db.NewSelect().Model(&propiedades).Order("id ASC").Scan(context.Background())
	return propiedades
}

func HandleHome(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, home.Home(fetchPropiedades(db.Create())))
}
