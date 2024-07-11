package db

import (
	"context"
	"fmt"

	"PagProp/models"
	"github.com/uptrace/bun"
)

func runMigrations(db *bun.DB) error {
	_, err := db.NewCreateTable().Model((*models.Propiedad)(nil)).IfNotExists().Exec(context.Background())
	if err != nil {
		return fmt.Errorf("error creating table: %w", err)
	}
	fmt.Println("Migrations completed successfully")
	return nil
}
