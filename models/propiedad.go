package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Propiedad struct {
	bun.BaseModel `bun:"table:propiedades,alias:p"`

	ID              int       `bun:"id,pk,autoincrement"`
	Comuna          string    `bun:"comuna,notnull"`
	Ciudad          string    `bun:"ciudad,notnull"`
	CalleNumero     string    `bun:"calle_numero,notnull"`
	Categoria       string    `bun:"categoria,notnull"`
	Venta           bool      `bun:"venta,notnull"`
	Arriendo        bool      `bun:"arriendo,notnull"`
	PrecioVenta     int       `bun:"precio_venta"`
	PrecioArriendo  int       `bun:"precio_arriendo"`
	NumFotos        int       `bun:"num_fotos"`
	FotosLinks      []string  `bun:"fotos_links,array"`
	TimeCreated     time.Time `bun:"time_created,notnull"`
	MetrosCuadrados int       `bun:"metros_cuadrados"`
	Dormitorios     int       `bun:"dormitorios"`
	Banios          int       `bun:"banios"`
	Descripcion     string    `bun:"descripcion,notnull"`
	GastosComunes   int       `bun:"gastos_comunes"`
}
