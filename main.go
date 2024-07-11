package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"PagProp/auth"

	"github.com/uptrace/bun"
	//"database/sql"

	"PagProp/db"
	"PagProp/handlers"

	"PagProp/models"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	database := db.Create()

	if err := database.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	if err := populateDatabase(database); err != nil {
		log.Fatalf("Failed to populate database: %v", err)
	}

	fmt.Printf("imprimiendo data")

	if err := queryAndPrintAllData(database); err != nil {
		log.Fatalf("Failed to query and print data: %v", err)
	}

	queryAndPrintAllData(database)

	router := chi.NewMux()

	router.Handle("/*", public())
	router.Get("/", handlers.Make(handlers.HandleHome))
	router.Get("/propiedades", handlers.Make(handlers.HandleHome))
	router.Get("/propiedades/{id}", handlers.Make(handlers.HandleDisplayPropiedad))
	router.Get("/search", handlers.Make(handlers.HandleSearch(database)))
	router.Get("/search-results", handlers.Make(handlers.HandleSearchResults(database)))

	router.Group(func(r chi.Router) {
		r.Use(auth.AuthMiddleware)
		r.Get("/uploadPropiedad", handlers.Make(handlers.HandleUploadPropiedadForm))
		r.Post("/uploadPropiedad", handlers.Make(handlers.HandleUploadPropiedad))
	})

	// Auth routes
	router.Get("/login", handlers.Make(handlers.HandleLoginForm))
	router.Post("/login", handlers.Make(handlers.HandleLogin))
	router.Get("/logout", handlers.Make(handlers.HandleLogout))

	listenAddr := os.Getenv("LISTEN_ADDR")
	slog.Info("HTTP server started", "listenAddr", listenAddr)
	http.ListenAndServe(listenAddr, router)

	fmt.Println("hello world")
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to db")
}

func queryAndPrintAllData(db *bun.DB) error {
	var propiedades []models.Propiedad
	err := db.NewSelect().Model(&propiedades).Scan(context.Background())
	if err != nil {
		return fmt.Errorf("error querying data: %w", err)
	}
	fmt.Printf("Total number of properties: %d\n", len(propiedades))

	return nil
}
