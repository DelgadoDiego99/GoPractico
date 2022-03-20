package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	inicio := time.Now()
	canal := make(chan string)

	servidores := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}

	for _, servidor := range servidores {
		go revisarServidor(servidor, canal)
	}

	for i := 0; i < len(servidores); i++ {
		fmt.Println(<-canal)
	}

	tiempoPasado := time.Since(inicio)
	fmt.Printf("Tiempo de ejecucción es %s", tiempoPasado)

}

func revisarServidor(servidor string, canal chan string) {
	_, err := http.Get(servidor)
	if err != nil {
		fmt.Println(servidor, "no está disponible")
		canal <- servidor + "no está disponible"
	} else {
		fmt.Println(servidor, "está funcionando correctamente")
		canal <- servidor + "está funcionando correctamente"
	}

}
