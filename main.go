package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
	Id    string `json:"Id"`
	Name  string `json:"Name"`
	Desc  string `json:"desc"`
	Price int    `json:"price"`
}

var Products []Product

func healthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Healthy!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllProducts")
	json.NewEncoder(w).Encode(Products)
}

func createNewProduct(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var product Product
	json.Unmarshal(reqBody, &product)
	Products = append(Products, product)

	json.NewEncoder(w).Encode(product)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, product := range Products {
		if product.Id == id {
			Products = append(Products[:index], Products[index+1:]...)
		}
	}

}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", healthcheck)
	myRouter.HandleFunc("/products", returnAllProducts)
	myRouter.HandleFunc("/product", createNewProduct).Methods("POST")
	myRouter.HandleFunc("/product/{id}", deleteProduct).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	Products = []Product{
		{Id: "1", Name: "Airpods Pro", Desc: "Apple Airpods Pro", Price: 350},
		{Id: "2", Name: "Cat Tower", Desc: "4 tier cat play tower", Price: 50},
	}
	handleRequests()
}
