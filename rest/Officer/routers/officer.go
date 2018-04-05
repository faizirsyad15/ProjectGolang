package routers

import (
	"day15/Officer/controllers"

	"github.com/gorilla/mux"
)

func setOfficerRouters(router *mux.Router) *mux.Router {
	// 1.b Membuat Fungsi di Controller Terlebih Dahulu
	router.HandleFunc("/officer", controllers.GetOfficer).Methods("GET")
	return router
}
