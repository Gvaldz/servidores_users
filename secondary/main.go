package main

import (
	"net/http"
	"io/ioutil"
	"time"
	"log"
)

var lastData []byte

func main() {
	go pollServerPrincipal()

	http.HandleFunc("/users", handleUsers)
	log.Println("Servidor secundario escuchando en :8081...")
	http.ListenAndServe(":8081", nil)
}

func pollServerPrincipal() {
	for {
		resp, err := http.Get("http://localhost:8080/users")
		if err != nil {
			log.Printf("Error al consultar el servidor principal: %v\n", err)
			time.Sleep(5 * time.Second) 
			continue
		}

		body, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			log.Printf("Error al leer la respuesta del servidor principal: %v\n", err)
			time.Sleep(5 * time.Second)
			continue
		}

		if string(body) != string(lastData) {
			log.Println("Datos actualizados en el servidor principal")
			lastData = body
		}

		time.Sleep(5 * time.Second)
	}
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(lastData)
}