package main

import (
	"encoding/json"
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

type MetaData interface{}

type User struct {
	Name  string `json:"name"` // `` permite que en el retorno, tenga la forma que se escribe
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func (u *User) ToJSON() ([]byte, error) {
	return json.Marshal(u)
}
