package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"github.com/PeaceOloruntoba/bookish-meme/internal/models"
	"github.com/jackc/pgx/v5"
)

type CarHandler struct {
	DB *pgx.Conn
}

func (h *CarHandler) GetCars(w http.ResponseWriter, r *http.Request) {
	rows, _ := h.DB.Query(context.Background(), "SELECT id, make, model, year FROM cars")
	var cars []models.Car
	
	for rows.Next() {
		var c models.Car
		rows.Scan(&c.ID, &c.Make, &c.Model, &c.Year)
		cars = append(cars, c)
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)
}