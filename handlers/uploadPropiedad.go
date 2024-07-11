package handlers

import (
	"net/http"
	"strconv"
	"time"

	"PagProp/db"
	"PagProp/models"
	"PagProp/views/uploadPropiedad"
)

func HandleUploadPropiedadForm(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, uploadPropiedad.UploadForm())
}

func HandleUploadPropiedad(w http.ResponseWriter, r *http.Request) error {
	// Parse the form data
	err := r.ParseForm()
	if err != nil {
		return err
	}

	precioVenta, _ := strconv.Atoi(r.FormValue("precio_venta"))
	precioArriendo, _ := strconv.Atoi(r.FormValue("precio_arriendo"))
	numFotos, _ := strconv.Atoi(r.FormValue("num_fotos"))
	metrosCuadrados, _ := strconv.Atoi(r.FormValue("metros_cuadrados"))
	dormitorios, _ := strconv.Atoi(r.FormValue("dormitorios"))
	banios, _ := strconv.Atoi(r.FormValue("banios"))
	gastosComunes, _ := strconv.Atoi(r.FormValue("gastos_comunes"))

	venta := r.FormValue("venta") == "on"
	arriendo := r.FormValue("arriendo") == "on"

	fotosLinks := r.Form["fotos_links[]"]

	propiedad := &models.Propiedad{
		Comuna:          r.FormValue("comuna"),
		Ciudad:          r.FormValue("ciudad"),
		CalleNumero:     r.FormValue("calle_numero"),
		Categoria:       r.FormValue("categoria"),
		Venta:           venta,
		Arriendo:        arriendo,
		PrecioVenta:     precioVenta,
		PrecioArriendo:  precioArriendo,
		NumFotos:        numFotos,
		FotosLinks:      fotosLinks,
		TimeCreated:     time.Now(),
		MetrosCuadrados: metrosCuadrados,
		Dormitorios:     dormitorios,
		Banios:          banios,
		Descripcion:     r.FormValue("descripcion"),
		GastosComunes:   gastosComunes,
	}
	// Insert the new property into the database
	database := db.Create()
	_, err = database.NewInsert().Model(propiedad).Exec(r.Context())
	if err != nil {
		return err
	}

	// Redirect to the home page or a success page
	http.Redirect(w, r, "/", http.StatusSeeOther)
	return nil
}
