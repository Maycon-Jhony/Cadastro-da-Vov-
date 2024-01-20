package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Cliente struct {
	ID       string `json:"id"`
	Nome     string `json:"nome"`
	Email    string `json:"email"`
	Telefone string `json:"telefone"`
}

var clientes []Cliente

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/clientes", GetClientes).Methods("GET")
	router.HandleFunc("/clientes/{id}", GetCliente).Methods("GET")
	router.HandleFunc("/clientes", CreateCliente).Methods("POST")
	router.HandleFunc("/clientes/{id}", UpdateCliente).Methods("PUT")
	router.HandleFunc("/clientes/{id}", DeleteCliente).Methods("DELETE")

	fmt.Println("Servidor escutando na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func GetClientes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clientes)
}

func GetCliente(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, cliente := range clientes {
		if cliente.ID == params["id"] {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(cliente)
			return
		}
	}
	http.NotFound(w, r)
}

func CreateCliente(w http.ResponseWriter, r *http.Request) {
	var cliente Cliente
	_ = json.NewDecoder(r.Body).Decode(&cliente)

	cliente.ID = fmt.Sprintf("%d", len(clientes)+1)
	clientes = append(clientes, cliente)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cliente)
}

func UpdateCliente(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var clienteAtualizado Cliente
	_ = json.NewDecoder(r.Body).Decode(&clienteAtualizado)

	for i, cliente := range clientes {
		if cliente.ID == params["id"] {
			clienteAtualizado.ID = cliente.ID
			clientes[i] = clienteAtualizado

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(clienteAtualizado)
			return
		}
	}
	http.NotFound(w, r)
}

func DeleteCliente(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, cliente := range clientes {
		if cliente.ID == params["id"] {
			clientes = append(clientes[:i], clientes[i+1:]...)

			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, "Cliente com o ID %s foi removido", params["id"])
			return
		}
	}
	http.NotFound(w, r)
}
