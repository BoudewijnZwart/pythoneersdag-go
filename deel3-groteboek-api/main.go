package main

import (
	"log"
	"net/http"

	"groteboek-api/internal/database"
	"groteboek-api/internal/handler"
	"groteboek-api/internal/repository"
)

func main() {
	db, err := database.New("groteboek.db")
	if err != nil {
		log.Fatalf("kon database niet openen: %v", err)
	}
	defer db.Close()

	cadeauRepository := repository.NewCadeauRepository(db)
	kindRepository := repository.NewKindRepository(db, cadeauRepository)
	kindHandler := handler.NewKindHandler(kindRepository)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /kinderen", kindHandler.List)
	mux.HandleFunc("GET /kinderen/{naam}", kindHandler.Get)
	mux.HandleFunc("POST /kinderen", kindHandler.Create)

	log.Println("Het grote boek van Sinterklaas luistert op :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
