package auth

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
)

func (i *AuthAPI) Login(w http.ResponseWriter, r *http.Request) {
	var request requestLogin
	json.NewDecoder(r.Body).Decode(&request)
	// Проверяем что пришел именно гуид
	_, err := uuid.Parse(request.Id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("id is not a valid UUID"))
		return
	}
	atoken, rtoken, err := i.authService.Create(r.Context(), request.Id)
	if err != nil {
		handleError(w, err)
		return
	}
	responseJson, err := json.Marshal(response{AccessToken: atoken, RefreshToken: rtoken})
	if err != nil {
		handleError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJson)
}

type requestLogin struct {
	Id string `json:"id"`
}
