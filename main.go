package main

import (
	"encoding/json"
	"log"

	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	ID string `json:"id,omitempty"`

	CityOrigin string `json:"origin,omitempty"`

	CityDestiny string `json:"destiny,omitempty"`

	Coordinate *Coordinate `json:"coordinate,omitempty"`
}

type Coordinate struct {
	LatOri string `json:"lat_ori,omitempty"`

	LongOri string `json:"long_ori,omitempty"`

	LatDest string `json:"lat_dest,omitempty"`

	LongDest string `json:"long_dest,omitempty"`
}

var routes []Route

// função principal

func main() {

	routes = append(routes, Route{ID: "1", CityOrigin: "São Paulo", CityDestiny: "Belo Horizonte", Coordinate: &Coordinate{LatOri: "-23.6821604", LongOri: "-46.6683344", LatDest: "-19.9026615", LongDest: "-44.1041375"}})

	routes = append(routes, Route{ID: "2", CityOrigin: "Rio de Janeiro", CityDestiny: "Salvador", Coordinate: &Coordinate{LatOri: "-22.913885", LongOri: "-43.7261792", LatDest: "-12.8754343", LongDest: "-38.6417385"}})

	routes = append(routes, Route{ID: "3", CityOrigin: "Curitiba", CityDestiny: "´Vitória"})

	router := mux.NewRouter()
	router.HandleFunc("/route", GetRoutes).Methods("GET")
	router.HandleFunc("/route/{id}", GetRoute).Methods("GET")
	router.HandleFunc("/route/{id}", CreateRoute).Methods("POST")
	router.HandleFunc("/route/{id}", DeleteRoute).Methods("DELETE")
	log.Fatal(http.ListenAndServe("localhost:3000", router))

}

func GetRoutes(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(routes)
}

func GetRoute(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for _, item := range routes {

		if item.ID ==
			params["id"] {

			json.NewEncoder(w).Encode(item)

			return

		}

	}

	json.NewEncoder(w).Encode(&Route{})

}

func CreateRoute(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var person Route

	_ =
		json.NewDecoder(r.Body).Decode(&person)

	person.ID = params["id"]

	routes = append(routes, person)

	json.NewEncoder(w).Encode(routes)
}

func DeleteRoute(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range routes {
		if item.ID == params["id"] {
			routes = append(routes[:index], routes[:index+1]...)
			break
		}
		json.NewEncoder(w).Encode(routes)

	}
}
