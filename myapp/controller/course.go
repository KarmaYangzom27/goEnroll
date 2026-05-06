package controller

import (
	"database/sql"
	"encoding/json"
	"myapp/model"
	"myapp/utils/httpResp"
	"net/http"

	"github.com/gorilla/mux"
)

func AddCourse(w http.ResponseWriter, r *http.Request) {
	var c model.Course
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&c); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "invalid json body")
		return
	}
	defer r.Body.Close()

	if err := c.Create(); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusCreated, map[string]string{"status": "course added"})
}

func GetCourse(w http.ResponseWriter, r *http.Request) {
	cid := mux.Vars(r)["cid"]
	c := model.Course{CId: cid}
	if err := c.Read(); err != nil {
		switch err {
		case sql.ErrNoRows:
			httpResp.RespondWithError(w, http.StatusNotFound, "course not found")
		default:
			httpResp.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, c)
}

func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	oldCid := mux.Vars(r)["cid"]
	var c model.Course
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&c); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "invalid json body")
		return
	}
	defer r.Body.Close()

	if err := c.Update(oldCid); err != nil {
		httpResp.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, c)
}

func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	cid := mux.Vars(r)["cid"]
	c := model.Course{CId: cid}
	if err := c.Delete(); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "deleted"})
}

func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	courses, err := model.GetAllCourses()
	if err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, courses)
}
