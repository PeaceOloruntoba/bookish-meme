package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/joho/godotenv"
	"github.com/PeaceOloruntoba/bookish-meme/internal/database"
	"github.com/PeaceOloruntoba/bookish-meme/internal/handlers"
)

func main() {
	godotenv.Load() // Load .env file
	
	conn, err := database.ConnectDB()
	if err != nil {
		log.Fatal("Could not connect to DB")
	}
	defer conn.Close(context.Background())

	h := &handlers.CarHandler{DB: conn}

	http.HandleFunc("GET /cars", h.GetCars)
    // Add other routes here...

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}