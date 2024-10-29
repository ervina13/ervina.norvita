package BukuController

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/config"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/config/helper"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/models/BukuModels"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/responses"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/responses/formaterror"
	"github.com/biezhi/gorm-paginator/pagination"
)

func LihatJenisBuku(w http.ResponseWriter, r *http.Request) {
	buku := BukuModels.JenisBuku{}
	bukus, err := buku.LihatJenisBuku(config.Db)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	query := r.URL.Query()
	id, _ := strconv.ParseUint(query.Get("id"), 10, 32)
	// Lihat Detail jenis buku
	if id != 0 {
		bukus, err := buku.DetailJenisBuku(config.Db, uint32(id))
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		// fmt.Println(bukus)
		responses.JSON(w, http.StatusOK, responses.Sukses(bukus))
		return
	}

	if query.Get("q") != "" {
		// fmt.Println(query.Get("q"))
		_, bukus := buku.CariJenisBuku(config.Db, query.Get("q"))
		var buku []BukuModels.JenisBuku
		paginator := pagination.Paging(&pagination.Param{
			DB:    bukus,
			Page:  int(helper.StringkeInt(query.Get("page"))),
			Limit: 3,
			// OrderBy: []string{"id desc"},
		}, &buku)
		responses.JSON(w, http.StatusOK, responses.Sukses(paginator))
		return
	}
	paginator := pagination.Paging(&pagination.Param{
		DB:    config.Db,
		Page:  int(helper.StringkeInt(query.Get("page"))),
		Limit: 3,
		// OrderBy: []string{"id desc"},
	}, &bukus)
	responses.JSON(w, http.StatusOK, responses.Sukses(paginator))
}

func TambahJenisBuku(w http.ResponseWriter, r *http.Request) {
	var jenbuk BukuModels.JenisBuku
	// decode data json request ke struct jenbuk
	err := json.NewDecoder(r.Body).Decode(&jenbuk)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	err = jenbuk.Validasi()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	// fmt.Print(jenbuk)
	jenisCreated, err := jenbuk.TambahJenisBuku(config.Db)

	if err != nil {

		formattedError := formaterror.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, jenisCreated.IDJenis))
	responses.JSON(w, http.StatusCreated, jenisCreated)
}

func HapusJenisBuku(w http.ResponseWriter, r *http.Request) {
	var jenbuk BukuModels.JenisBuku
	query := r.URL.Query()
	id, _ := strconv.ParseUint(query.Get("id"), 10, 32)
	if id == 0 {
		responses.ERROR(w, http.StatusInternalServerError, fmt.Errorf("URL Invalid"))
		return
	}
	_, err := jenbuk.HapusJenisBuku(config.Db, uint32(id))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	// w.Header().Set("Entity", fmt.Sprintf("%d", id))
	responses.JSON(w, http.StatusOK, responses.Sukses(fmt.Sprintf("Data dengan id %d sukses terhapus ", id)))
}

func UpdateJenisBuku(w http.ResponseWriter, r *http.Request) {
	var jenbuk BukuModels.JenisBuku
	// decode data json request ke struct jenbuk
	query := r.URL.Query()
	id, _ := strconv.ParseUint(query.Get("id"), 10, 32)
	if id == 0 {
		responses.ERROR(w, http.StatusInternalServerError, fmt.Errorf("URL Invalid"))
		return
	}
	err := json.NewDecoder(r.Body).Decode(&jenbuk)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	err = jenbuk.Validasi()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	updated, err := jenbuk.UpdateJenisBuku(config.Db, uint32(id))
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, responses.Sukses(updated))
	// w.Header().Set("Entity", fmt.Sprintf("%d", id))
	// responses.JSON(w, http.StatusOK, responses.Sukses(fmt.Sprintf("Data dengan id %d sukses terhapus ", id)))
}
