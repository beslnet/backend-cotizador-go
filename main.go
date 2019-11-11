package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type product struct {
	Id      string `json:"id"`
	Code    string `json:"Code"`
	Name    string `json:"Name"`
	Country string `json:"Country"`
	Price   string `json:"Price"`
	Type    string `json:"Type"`
}

type allProducts []product

var products = allProducts{
	{
		Id:      "1",
		Code:    "Cherry",
		Name:    "BLIN CHERRY",
		Country: "USA",
		Price:   "50.00",
		Type:    "DOLAR",
	},
	{
		Id:      "2",
		Code:    "Cherry",
		Name:    "BLIN CHERRY",
		Country: "SPAIN",
		Price:   "55.00",
		Type:    "EURO",
	},
	{
		Id:      "3",
		Code:    "STP",
		Name:    "BLIN STEAMPUNK",
		Country: "FRANCE",
		Price:   "18.00",
		Type:    "EURO",
	},
	{
		Id:      "4",
		Code:    "STP",
		Name:    "BLIN STEAMPUNK",
		Country: "TAIWAN",
		Price:   "15.00",
		Type:    "DOLAR",
	},
	{
		Id:      "5",
		Code:    "Cherry",
		Name:    "FLAMINGO WARHOL",
		Country: "USA",
		Price:   "10.00",
		Type:    "DOLAR",
	},
	{
		Id:      "6",
		Code:    "Cherry",
		Name:    "FLAMINGO WARHOL",
		Country: "SPAIN",
		Price:   "8.50",
		Type:    "EURO",
	},
}

func getOneProduct(w http.ResponseWriter, r *http.Request) {

	cors(w, r)

	name := mux.Vars(r)["Name"]
	country := mux.Vars(r)["Country"]

	for _, singleProduct := range products {
		if singleProduct.Name == name && singleProduct.Country == country {
			json.NewEncoder(w).Encode(singleProduct)
		}
	}

}

func getAllProducts(w http.ResponseWriter, r *http.Request) {

	cors(w, r)

	json.NewEncoder(w).Encode(products)
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/products", getAllProducts).Methods("GET")
	router.HandleFunc("/getProduct/{Name}/{Country}", getOneProduct).Methods("GET")
	log.Fatal(http.ListenAndServe(":9090", router))
}

func cors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
}
