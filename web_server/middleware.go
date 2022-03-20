package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Los middleware son algo como minihandlers pero reciclables

func CheckAuth() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			flag := true
			fmt.Println("Autenticando...")
			if flag {
				f(w, r)
			} else {
				return
			}
		}
	}
}

func Logging() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() { // Funcion anónima, porque es la única vez que se va a usar
				log.Println(r.URL.Path, time.Since(start)) // Log imprime la hora del computador más lo parámetros que se agreguen
			}()
			f(w, r)
		}
	}
}
