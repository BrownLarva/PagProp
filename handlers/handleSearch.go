package handlers

import (
	"context"
	"net/http"
	"strconv"

	"PagProp/models"
	"PagProp/views/search"

	"github.com/uptrace/bun"
)

func fetchComunas(db *bun.DB) ([]string, error) {
	var comunas []string
	err := db.NewSelect().
		Model((*models.Propiedad)(nil)).
		ColumnExpr("DISTINCT comuna").
		Order("comuna ASC").
		Scan(context.Background(), &comunas)
	return comunas, err
}

func fetchMaxPrice(db *bun.DB) (int, error) {
	var maxPrice int
	err := db.NewSelect().
		Model((*models.Propiedad)(nil)).
		ColumnExpr("GREATEST(MAX(precio_venta), MAX(precio_arriendo)) AS max_price").
		Scan(context.Background(), &maxPrice)
	return maxPrice, err
}

func HandleSearch(db *bun.DB) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		comunas, err := fetchComunas(db)
		if err != nil {
			return err
		}

		maxPrice, err := fetchMaxPrice(db)
		if err != nil {
			return err
		}

		var results []models.Propiedad
		if r.URL.Query().Get("comuna") != "" || r.URL.Query().Get("categoria") != "" ||
			r.URL.Query().Get("min_price") != "" || r.URL.Query().Get("max_price") != "" {
			results, err = performSearch(db, r)
			if err != nil {
				return err
			}
		}

		return Render(w, r, search.SearchPage(comunas, maxPrice, results))
	}
}

func performSearch(db *bun.DB, r *http.Request) ([]models.Propiedad, error) {
	comuna := r.URL.Query().Get("comuna")
	categoria := r.URL.Query().Get("categoria")
	minPrice, _ := strconv.Atoi(r.URL.Query().Get("min_price"))
	maxPrice, _ := strconv.Atoi(r.URL.Query().Get("max_price"))

	var propiedades []models.Propiedad
	query := db.NewSelect().Model(&propiedades)

	if comuna != "" {
		query = query.Where("comuna = ?", comuna)
	}
	if categoria != "" {
		query = query.Where("categoria = ?", categoria)
	}
	if minPrice > 0 {
		query = query.Where("precio_venta >= ? OR precio_arriendo >= ?", minPrice, minPrice)
	}
	if maxPrice > 0 {
		query = query.Where("precio_venta <= ? OR precio_arriendo <= ?", maxPrice, maxPrice)
	}

	err := query.Scan(r.Context())
	return propiedades, err
}

func HandleSearchResults(db *bun.DB) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		results, err := performSearch(db, r)
		if err != nil {
			return err
		}
		return Render(w, r, search.SearchResults(results))
	}
}
