package main

import (
	"fmt"
	"lutadores-modulo/configs"
	"lutadores-modulo/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Post("/", handlers.Create)
	r.Get("/", handlers.List)
	r.Get("/{id}", handlers.Get)
	r.Delete("/{id}", handlers.Delete)
	r.Put("/{id}", handlers.Update)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)

}
