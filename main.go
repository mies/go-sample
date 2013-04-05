package main

import (
	"os"
	"log"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, req *http.Request) {
	cities, _ := json.Marshal("{'Amsterdam', 'San Francisco' 'Berlin', 'New York'}")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(cities)
}

func routerHandler(r *mux.Router) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		r.ServeHTTP(w, req)
	}
}

func router() * mux.Router{
	router := mux.NewRouter()
	router.HandleFunc("/cities.json", index).Methods("GET")
	return router
}

func main() {
	handler := routerHandler(router())
	err := http.ListenAndServe(":"+os.Getenv("PORT"), handler)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
