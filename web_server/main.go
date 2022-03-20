// Usar go run . para ejecutar todos los archivos de la carpeta

package main

func main() {

	server := NewServer("3000")
	server.Listen()

}
