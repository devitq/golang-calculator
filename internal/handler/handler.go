package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"calc_service/internal/calculator"
)

type CalculateRequest struct {
	Expression string `json:"expression"`
}

type CalculateResponse struct {
	Result *float64 `json:"result,omitempty"`
	Error  *string  `json:"error,omitempty"`
}

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, `{"error":"Method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	var req CalculateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"Invalid request format"}`, http.StatusBadRequest)
		return
	}

	result, err := calculator.Calc(req.Expression)
	if err != nil {
		if calculator.IsValidationError(err) {
			http.Error(w, fmt.Sprintf(`{"error":"Expression is not valid: %s"}`, err.Error()), http.StatusUnprocessableEntity)
		} else {
			http.Error(w, `{"error":"Internal server error"}`, http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(CalculateResponse{Result: &result})
}
