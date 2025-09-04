package menfesscontroller

import (
	"net/http"
	"encoding/json"
)

func (mc *MenfessController) Login(w http.ResponseWriter, r *http.Request) {
	var loginReq entities.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	
}