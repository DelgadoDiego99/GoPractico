// Usar go run . para ejecutar todos los archivos de la carpeta

package main

func main() {

	server := NewServer("3000")
	server.Handle("/", "GET", HandleRoot)
	server.Handle("/api", "POST", server.AddMiddleware(HandleHome, CheckAuth(), Logging()))
	//server.Handle("/api", "GET", HandleHome)
	server.Listen()

}
