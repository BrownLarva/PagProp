package main

import (
	"context"
	"fmt"
	"time"

	"PagProp/models"
	"github.com/uptrace/bun"
)

func populateDatabase(db *bun.DB) error {
	// Check the number of existing entries

	//_, err := db.NewDelete().Model((*models.Propiedad)(nil)).Exec(context.Background())
	//if err != nil {
	//	return fmt.Errorf("error deleting existing properties: %v", err)
	//}

	count, err := db.NewSelect().Model((*models.Propiedad)(nil)).Count(context.Background())
	if err != nil {
		return fmt.Errorf("error counting existing properties: %v", err)
	}

	if count >= 5 {
		fmt.Println("Database already has 5 or more entries. Skipping population.")
		return nil
	}
	properties := []models.Propiedad{
		{
			Comuna: "Concon", Ciudad: "Bosques de Montemar", CalleNumero: "Avenida del Mar 123", Categoria: "Casa",
			Venta: true, Arriendo: false, PrecioVenta: 250000000, NumFotos: 1,
			FotosLinks: []string{"https://www.enaco.cl/img/sitemap/comprar-casa-departamento.jpg"}, TimeCreated: time.Now(), MetrosCuadrados: 150,
			Dormitorios: 3, Banios: 2, Descripcion: "Hermosa casa con vista al mar", GastosComunes: 100000,
		},
		{
			Comuna: "Viña del Mar", Ciudad: "Viña Centro", CalleNumero: "Calle Valparaíso 456", Categoria: "Departamento",
			Venta: false, Arriendo: true, PrecioArriendo: 800000, NumFotos: 1,
			FotosLinks: []string{"https://rastro.com/fotos3/2017/12/12/2255t73udgg6g2amgsnhdk99h1_2017121205264515a303b853f6df.jpg"}, TimeCreated: time.Now(), MetrosCuadrados: 80,
			Dormitorios: 2, Banios: 1, Descripcion: "Céntrico departamento cerca de la playa", GastosComunes: 80000,
		},
		{
			Comuna: "Concon", Ciudad: "Bosques de Montemar", CalleNumero: "Camino a las Dunas 789", Categoria: "Casa",
			Venta: true, Arriendo: true, PrecioVenta: 300000000, PrecioArriendo: 1200000, NumFotos: 1,
			FotosLinks: []string{"https://www.burgospro.cl/wp-content/uploads/2021/09/IMG_4385.jpg"}, TimeCreated: time.Now(), MetrosCuadrados: 200,
			Dormitorios: 4, Banios: 3, Descripcion: "Amplia casa familiar con jardín", GastosComunes: 150000,
		},
		{
			Comuna: "Viña del Mar", Ciudad: "Viña Centro", CalleNumero: "Avenida Libertad 101", Categoria: "Departamento",
			Venta: true, Arriendo: false, PrecioVenta: 180000000, NumFotos: 1,
			FotosLinks: []string{"https://imanquehue.com/content/uploads/departamento-piloto.png"}, TimeCreated: time.Now(), MetrosCuadrados: 70,
			Dormitorios: 2, Banios: 1, Descripcion: "Moderno departamento en el corazón de Viña", GastosComunes: 70000,
		},
		{
			Comuna: "Concon", Ciudad: "Bosques de Montemar", CalleNumero: "Paseo de los Pinos 202", Categoria: "Casa",
			Venta: false, Arriendo: true, PrecioArriendo: 1500000, NumFotos: 1,
			FotosLinks: []string{"https://images.adsttc.com/media/images/5eaa/e8ea/b357/653d/aa00/0269/medium_jpg/_fi.jpg?1588259042"}, TimeCreated: time.Now(), MetrosCuadrados: 180,
			Dormitorios: 3, Banios: 2, Descripcion: "Casa de lujo con piscina", GastosComunes: 200000,
		},
		{
			Comuna: "Viña del Mar", Ciudad: "Viña Centro", CalleNumero: "Calle Quillota 303", Categoria: "Departamento",
			Venta: true, Arriendo: true, PrecioVenta: 150000000, PrecioArriendo: 700000, NumFotos: 1,
			FotosLinks: []string{"https://www.laflorida.cl/sitio/wp-content/uploads/2022/11/WhatsApp-Image-2022-11-11-at-12.01.47.jpeg"}, TimeCreated: time.Now(), MetrosCuadrados: 90,
			Dormitorios: 3, Banios: 2, Descripcion: "Acogedor departamento con vista a la ciudad", GastosComunes: 90000,
		},
		{
			Comuna: "Concon", Ciudad: "Bosques de Montemar", CalleNumero: "Avenida Borgoño 404", Categoria: "Casa",
			Venta: true, Arriendo: false, PrecioVenta: 350000000, NumFotos: 1,
			FotosLinks: []string{"https://idea.cl/web/wp-content/uploads/2021/11/home-psm2.jpg"}, TimeCreated: time.Now(), MetrosCuadrados: 220,
			Dormitorios: 4, Banios: 3, Descripcion: "Espectacular casa con terraza y vista panorámica", GastosComunes: 180000,
		},
		{
			Comuna: "Viña del Mar", Ciudad: "Viña Centro", CalleNumero: "Calle Álvarez 505", Categoria: "Departamento",
			Venta: false, Arriendo: true, PrecioArriendo: 600000, NumFotos: 1,
			FotosLinks: []string{"https://img-us-1.trovit.com/img1mx/N1L_1dxWV9/N1L_1dxWV9.5_11.jpg"}, TimeCreated: time.Now(), MetrosCuadrados: 60,
			Dormitorios: 1, Banios: 1, Descripcion: "Estudio perfecto para estudiantes o profesionales", GastosComunes: 50000,
		},
		{
			Comuna: "Concon", Ciudad: "Bosques de Montemar", CalleNumero: "Calle del Bosque 606", Categoria: "Casa",
			Venta: true, Arriendo: true, PrecioVenta: 280000000, PrecioArriendo: 1100000, NumFotos: 1,
			FotosLinks: []string{"https://img-us-1.trovit.com/img1mx/1yvzn1_1O181Q/1yvzn1_1O181Q.2_11.jpg"}, TimeCreated: time.Now(), MetrosCuadrados: 170,
			Dormitorios: 3, Banios: 2, Descripcion: "Casa familiar en entorno natural", GastosComunes: 130000,
		},
		{
			Comuna: "Viña del Mar", Ciudad: "Viña Centro", CalleNumero: "Avenida San Martín 707", Categoria: "Departamento",
			Venta: true, Arriendo: false, PrecioVenta: 200000000, NumFotos: 1,
			FotosLinks: []string{"https://photos.encuentra24.com/t_or_fh_m/f_auto/v1/cl/28/28/50/91/28285091_043a65"}, TimeCreated: time.Now(), MetrosCuadrados: 100,
			Dormitorios: 3, Banios: 2, Descripcion: "Amplio departamento a pasos de la playa", GastosComunes: 100000,
		},
	}

	// Insert properties into the database
	for _, prop := range properties {
		_, err := db.NewInsert().Model(&prop).Exec(context.Background())
		if err != nil {
			return fmt.Errorf("error inserting property: %v", err)
		}
	}

	fmt.Println("10 dummy properties inserted successfully!")
	return nil
}
