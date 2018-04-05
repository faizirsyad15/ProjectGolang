package controllers

import (
	"day15/Item/data"
	"encoding/json"
	"net/http"
)

func GetItem(w http.ResponseWriter, r *http.Request) {
	// ambil datanya
	// untuk ambil data memerlukan koneksi
	// 1.c buat koneksi
	context := Context{}
	d := DBInitial(context.DB)
	defer d.Close()
	// ambil data dari repositories
	// buat perintah di repositories
	// 1.d buat repo petugas
	repo := &data.ItemRepository{d}
	item := data.GetAll(repo)
	// olah error nya
	// tampilkan datanya
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//w.Write(data)
	mdl, _ := json.Marshal(item)
	w.Write(mdl)
}
