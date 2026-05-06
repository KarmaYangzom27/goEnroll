package controller

import (
	"encoding/json"
	"myapp/model"
	"myapp/utils/httpResp"
	"net/http"
	"time"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	var admin model.Admin
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&admin); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "invalid json body")
		return
	}
	defer r.Body.Close()

	saveErr := admin.Create()
	if saveErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, saveErr.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusCreated, map[string]string{"status": "admin added"})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var admin model.Admin
	err := json.NewDecoder(r.Body).Decode(&admin)
	if err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "invalid json body")
		return
	}
	defer r.Body.Close()

	getErr := admin.Get()
	if getErr != nil {
		httpResp.RespondWithError(w, http.StatusUnauthorized, getErr.Error())
		return
	}

	// set cookie
	cookie := http.Cookie{
		Name:    "my-cookie",
		Value:   "my-value",
		Expires: time.Now().Add(30 * time.Minute),
		Secure:  true,
	}
	http.SetCookie(w, &cookie)

	httpResp.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "login success"})
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "my-cookie",
		Expires: time.Now(),
	})
	httpResp.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "cookie deleted"})
}
