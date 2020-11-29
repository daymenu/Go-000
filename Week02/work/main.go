package main

import (
	"log"
	"net/http"

	"github.com/daymenu/Go-000/Week02/work/controller"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Panicln(err)
		}
	}()
	mux := http.DefaultServeMux
	mux.HandleFunc("/", controller.SearchMoive)
	mux.HandleFunc("/search/moive", controller.SearchMoive)
	log.Fatal(http.ListenAndServe(":9090", mux))
}
