package main

import (
	"encoding/json"
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

type User struct {
	Name  string `json:"name"` // <- Cuando se serialize en JSON, las llaves deben mandarse con este nombre
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type MetaData interface{}

func (u *User) ToJSON() ([]byte, error) {
	return json.Marshal(u)
}
