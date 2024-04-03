package routes

import (
	"encoding/json"
	"go-snowflake/database"
	"net/http"
)

type QueryRequest struct {
	Query string `json:"query"`
}

func QueryHandler(w http.ResponseWriter, r *http.Request) {
	var req QueryRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := database.Connect()
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	results, err := database.Query(db, req.Query)
	if err != nil {
		http.Error(w, "Failed to Query database", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
