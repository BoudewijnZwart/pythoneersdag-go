package main

import (
	"log"
	"net/http"

	"sinterklaas-api/db"
	"sinterklaas-api/handlers"
	"sinterklaas-api/repository"
)

func main() {
	database, err := db.New("sinterklaas.db")
	if err != nil {
		log.Fatalf("kon database niet openen: %v", err)
	}

	kindRepo := repository.NewKindRepository(database)
	cadeauRepo := repository.NewCadeauRepository(database)
	kindHandler := handlers.NewKindHandler(kindRepo, cadeauRepo)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /kinderen", kindHandler.List)
	mux.HandleFunc("GET /kinderen/{naam}", kindHandler.GetByNaam)
	mux.HandleFunc("POST /kinderen", kindHandler.Create)

	log.Println("server luistert op :8081")
	if err := http.ListenAndServe(":8081", mux); err != nil {
		log.Fatal(err)
	}
}
