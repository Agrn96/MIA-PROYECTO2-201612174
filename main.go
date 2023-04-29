package main

import (
	"MIA-PROYECTO2-201612174/interpreter"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	_ "github.com/godror/godror"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type dataJson struct { // structura para data recibida
	Entrada string `json:"entrada"`
}

// Funcionalidad del interprete
func interpreterF(input string, consola *string) {
	input = strings.ReplaceAll(input, ">", "-")
	//input = input + "\n"
	//fmt.Println(input)
	interpreter.CommandChecker(interpreter.ScanInput(input), consola)

}

func interpretarBackend(w http.ResponseWriter, r *http.Request) {
	var inicio dataJson
	var consola = ""
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error datos no validos")
	}
	json.Unmarshal(reqBody, &inicio)
	interpreterF(inicio.Entrada, &consola)
	inicio.Entrada = consola
	fmt.Println(consola)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(inicio)
}

func getHola(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("HOLA MUNDO")
}

// Funcion Main
func main() {
	router := mux.NewRouter().StrictSlash(false)
	headers := handlers.AllowedHeaders([]string{"X-Request-Width", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST"})
	origins := handlers.AllowedOrigins([]string{"*"})
	router.HandleFunc("/", getHola)
	router.HandleFunc("/interpretar", interpretarBackend).Methods("POST")
	log.Println("Listening at port 8080")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(router)))
}
