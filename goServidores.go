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

	//i := 0
	//for i < 3 { // for = while || for condicion = do while
	for _, servidor := range servidores {
		go revisarServidor(servidor, canal)
	}
	/*time.Sleep(1 * time.Second)
	fmt.Println(<-canal)
	i++*/
	//}

	// Imprime la cantidad de veces que sea igual al tamaño del slice de servidores
	/*for i := 0; i < len(servidores); i++ {
		fmt.Println(<-canal)
	}*/

	tiempoPasado := time.Since(inicio)
	fmt.Printf("Tiempo de ejecucción es %s", tiempoPasado)

}

func revisarServidor(servidor string, canal chan string) {
	_, err := http.Get(servidor)
	if err != nil {
		canal <- servidor + " no está disponible"
	} else {
		canal <- servidor + " está funcionando correctamente"
	}

}
