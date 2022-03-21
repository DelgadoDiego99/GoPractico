package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo")
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Estoy en el endpoint de API")
}

func PostRequest(w http.ResponseWriter, r *http.Request) {
	decodificador := json.NewDecoder(r.Body)
	var metaData MetaData
	err := decodificador.Decode(&metaData)
	if err != nil {
		fmt.Fprintf(w, "Error: %v || EOF: Error of file", err)
		return
	}
	fmt.Fprintf(w, "Payload %v\n", metaData)
}

func UserPostRequest(w http.ResponseWriter, r *http.Request) {
	decodificador := json.NewDecoder(r.Body)
	var user User
	err := decodificador.Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "Error: %v || EOF: Error of file", err)
		return
	}
	response, err := user.ToJSON()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(response) // Write traduce los []byte || No hay necesidad de escribir un w.Header() porque Write() detecta automáticamente el tipo

	//fmt.Fprintf(w, "Payload %v\n", user)
}
