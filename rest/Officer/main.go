package main

import (
	"day15/Officer/routers"
	"log"
	"net/http"
)

func main() {
	// 1. Buat Routing
	// Buat Fungsi Routing
	router := routers.InitRouters()

	// Buat Configurasi Server
	log.Fatal(http.ListenAndServe(":8882", router))
}
