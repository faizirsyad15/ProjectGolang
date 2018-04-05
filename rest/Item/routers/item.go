package routers

import (
	"day15/Item/controllers"

	"github.com/gorilla/mux"
)

func setItemRouters(router *mux.Router) *mux.Router {
	// 1.b Membuat Fungsi di Controller Terlebih Dahulu
	router.HandleFunc("/item", controllers.GetItem).Methods("GET")
	return router
}
