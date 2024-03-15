package auth

import (
	"encoding/json"
	"net/http"
)

func (i *AuthAPI) Refresh(w http.ResponseWriter, r *http.Request) {
	var request requestRefresh
	json.NewDecoder(r.Body).Decode(&request)
	atoken, rtoken, err := i.authService.Update(r.Context(), request.AccessToken, request.RefreshToken)
	if err != nil {
		handleError(w, err)
		return
	}
	responseJson, err := json.Marshal(responseRefresh{AccessToken: atoken, RefreshToken: rtoken})
	if err != nil {
		handleError(w, err)
		return
	}

	w.Write(responseJson)
	return
}

type requestRefresh struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
type responseRefresh struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
