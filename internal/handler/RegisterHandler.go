package handler

import (
	"encoding/json"
	"net/http"

	"server.com/auth-service/internal/dto"
)

func (h *AuthHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// validate the method
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", 405)
	}
	// Read the request
	var req dto.SignupRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	// validate the data
	if req.Email == "" || req.Password == "" || req.Name == "" {
		http.Error(w, "Not Allowed", http.StatusBadRequest)
	}
	// call the service
	user, e := h.service.SignupService(r.Context(), req.Name, req.Email, req.Password)
	// return the res
	if e != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(user)
}
