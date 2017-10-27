package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/JanoBotso/Tarea1-GG/server"
	"github.com/JanoBotso/Tarea1-GG/database"
)

func main() {
	db := database.New()


	myServer := server.New(db)

	router := mux.NewRouter()

	router.NewRoute().
		Path("/cities/").
		HandlerFunc(myServer.ListCiudades).
		Methods("GET").
		Name("ListCiudades")


	router.NewRoute().
		Path("/cities/").
		HandlerFunc(myServer.AddCiudades).
		Methods("POST").
		Name("AddCiudades")


	router.NewRoute().
		Path("/connection/s").
		HandlerFunc(myServer.ListConecciones).
		Methods("GET").
		Name("ListConecciones")


	router.NewRoute().
		Path("/connections/").
		HandlerFunc(myServer.AddConecciones).
		Methods("POST").
		Name("AddConecciones")

/*
	router.NewRoute().
		Path("/solve/").
		HandlerFunc(myServer.EntryGet).
		Methods("GET").
		Name("EntryGet")
*/

	http.ListenAndServe(":8000", router)
}
