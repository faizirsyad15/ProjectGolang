package main

import (
	"day15/Selling/routers"
	"log"
	"net/http"
)

func main() {
	router := routers.InitRouters()

	log.Fatal(http.ListenAndServe(":8883", router))
}
