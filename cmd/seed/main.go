package main

import (
	"context"
	"fmt"

	"PagProp/db"
	"PagProp/models"

	"github.com/uptrace/bun"
)

func main() {
	database := db.Create()

	populateData(database)

	queryAndPrintData(database)
}

func populateData(db *bun.DB) error {
	propiedades := []models.Propiedad{
		{Precio: 100000, Descripcion: "Casa grande con jardín", ImageLink: "https://example.com/casa1.jpg"},
		{Precio: 75000, Descripcion: "Apartamento céntrico", ImageLink: "https://example.com/apartamento1.jpg"},
		{Precio: 150000, Descripcion: "Chalet con piscina", ImageLink: "https://example.com/chalet1.jpg"},
	}

	_, err := db.NewInsert().Model(&propiedades).Exec(context.Background())
	if err != nil {
		return fmt.Errorf("error inserting data: %w", err)
	}
	fmt.Println("Dummy data inserted successfully")
	return nil
}

func queryAndPrintData(db *bun.DB) error {
	var propiedades []models.Propiedad
	err := db.NewSelect().Model(&propiedades).Scan(context.Background())
	if err != nil {
		return fmt.Errorf("error querying data: %w", err)
	}

	for _, p := range propiedades {
		fmt.Printf("ID: %d, Precio: %d, Descripción: %s, Image Link: %s\n", p.ID, p.Precio, p.Descripcion, p.ImageLink)
	}
	return nil
}
