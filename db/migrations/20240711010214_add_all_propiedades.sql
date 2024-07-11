-- +goose Up
DROP TABLE IF EXISTS propiedades;

CREATE TABLE propiedades (
	id SERIAL NOT NULL primary key,
	comuna          TEXT NOT NULL,
	ciudad          TEXT NOT NULL,
	calle_numero    TEXT NOT NULL,
	categoria       TEXT NOT NULL,
	venta           BOOLEAN NOT NULL,
	arriendo        BOOLEAN NOT NULL,
	precio_venta    INTEGER,
	precio_arriendo INTEGER,
	num_fotos       INTEGER,
	fotos_links     TEXT[] NOT NULL,
	time_created    TIMESTAMP NOT NULL,
	metros_cuadrados INTEGER,
	dormitorios     INTEGER,
	banios          INTEGER,
	descripcion     TEXT NOT NULL,
	gastos_comunes  INTEGER
);

-- +goose Down
DROP TABLE propiedades;
