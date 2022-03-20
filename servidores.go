package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	inicio := time.Now()

	servidores := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}

	for _, servidor := range servidores {
		revisarServidor(servidor)
	}
	tiempoPasado := time.Since(inicio)
	fmt.Printf("Tiempo de ejecucción es %s", tiempoPasado)

}

func revisarServidor(servidor string) {
	_, err := http.Get(servidor)
	if err != nil {
		fmt.Println(servidor, "no está disponible")
	} else {
		fmt.Println(servidor, "está funcionando correctamente")
	}

}
