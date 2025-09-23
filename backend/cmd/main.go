package main

import (
	"log"
	"net/http"

	"github.com/lucasmsaluno/my-tasks/internal/database"
	"github.com/lucasmsaluno/my-tasks/internal/repositories"
	"github.com/lucasmsaluno/my-tasks/internal/routes"
	"github.com/rs/cors"
)

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco: %v", err)
	}
	defer db.Close()

	repo := repositories.NewTaskRepository(db)
	router := routes.SetupRoutes(repo)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5500", "http://127.0.0.1:5500"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	log.Println("server open: http://localhost:8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
