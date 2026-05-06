package controller

import (
	"encoding/json"
	"myapp/model"
	"myapp/utils/httpResp"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func getDate() string {
	const apiDateLayout = "2006-01-02T15:04:05Z"
	date := time.Now().UTC()
	return date.Format(apiDateLayout)
}

func Enroll(w http.ResponseWriter, r *http.Request) {
	var e model.Enroll
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&e); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "invalid json body")
		return
	}
	e.Date_Enrolled = getDate()
	defer r.Body.Close()

	saveErr := e.EnrollStud()
	if saveErr != nil {
		if strings.Contains(saveErr.Error(), "duplicate key") {
			httpResp.RespondWithError(w, http.StatusForbidden, "Duplicate keys")
			return
		} else {
			httpResp.RespondWithError(w, http.StatusInternalServerError, saveErr.Error())
			return
		}
	}
	httpResp.RespondWithJSON(w, http.StatusCreated, map[string]string{"status": "enrolled"})
}

func GetEnroll(w http.ResponseWriter, r *http.Request) {
	sid := mux.Vars(r)["sid"]
	cid := mux.Vars(r)["cid"]
	stdid, _ := strconv.ParseInt(sid, 10, 64)

	e := model.Enroll{StdId: stdid, CourseID: cid}
	getErr := e.Get()
	if getErr != nil {
		httpResp.RespondWithError(w, http.StatusNotFound, "No such enrollments")
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, e)
}

func GetEnrolls(w http.ResponseWriter, r *http.Request) {
	enrolls, getErr := model.GetAllEnrolls()
	if getErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, getErr.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, enrolls)
}

func DeleteEnroll(w http.ResponseWriter, r *http.Request) {
	sid := mux.Vars(r)["sid"]
	cid := mux.Vars(r)["cid"]
	stdid, _ := strconv.ParseInt(sid, 10, 64)

	e := model.Enroll{StdId: stdid, CourseID: cid}
	if err := e.Delete(); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "deleted"})
}
