package routers

import (
	"day15/Selling/controllers"

	"github.com/gorilla/mux"
)

func setSellingRouters(router *mux.Router) *mux.Router {
	// 1.b Membuat Fungsi di Controller Terlebih Dahulu
	router.HandleFunc("/selling", controllers.GetSelling).Methods("GET")
	return router
}
