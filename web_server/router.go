package main

import (
	"net/http"
)

// Se hace mapeo doble, uno para los métodos y el otro para las rutas
type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

// Hay doble bool para retornar si existe el path, y el otro bool para retornar si el método existe en dicha ruta
func (r *Router) FindHandler(path, method string) (http.HandlerFunc, bool, bool) {
	_, exist := r.rules[path]
	handler, methodExist := r.rules[path][method] // Se escribe [path][method] porque URL solo hay una, por eso se escribe de primero y el mapeo se realiza de esa manera
	return handler, methodExist, exist
}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	handler, methodExist, exit := r.FindHandler(request.URL.Path, request.Method)
	if !exit {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if !methodExist {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	handler(w, request)
}
