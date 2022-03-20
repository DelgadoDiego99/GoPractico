package main

import "net/http"

type Server struct {
	port   string
	router *Router
}

// Devuelve el objeto real, copias no
func NewServer(port string) *Server {
	// Es & es para devolver la dirección de memoria donde se va a crear el objetos
	return &Server{
		port:   ":" + port,  // Agregar los ':' ahorra que cuando se cree un servidor, solo se tenga que colocar el número del puerto, y no ':#'
		router: NewRouter(), // Llama a la función de crear Router directamente

	}
}

func (s *Server) Handle(path, method string, handler http.HandlerFunc) {
	_, exist := s.router.rules[path]
	if !exist {
		s.router.rules[path] = make(map[string]http.HandlerFunc)
	}
	s.router.rules[path][method] = handler
}

// Se agrega ... cuando no se conoce la cantidad de parametros de ese tipo
func (s *Server) AddMiddleware(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func (s *Server) Listen() error {
	// El router va a ser el encargado se procesar las URL y procesarlas como se debe
	// El '/' represanta el punto de entrada y 's.router' el que va a manejar dicha ruta
	http.Handle("/", s.router)
	// 'ListenAndServe' escucha todas las peticiones y las resuelve
	// El parametro nil es para los handlers, porque no se quiere los handles predeterminados
	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		return err
	}
	return nil
}
