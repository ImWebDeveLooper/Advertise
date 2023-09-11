package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Data         interface{} `json:"data,omitempty"`
	Message      string      `json:"message,omitempty"`
	Error        error       `json:"-,omitempty"`
	ErrorMessage string      `json:"error_message,omitempty"`
}

func ErrorResponse(w http.ResponseWriter, resp Response) {
	jData, err := json.Marshal(resp)
	if err != nil {
		BadRequest(w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jData)
	if err != nil {
		return
	}
}
func BadRequest(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
}

func InternalServerError(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
}

func CreateResponse(w http.ResponseWriter, resp Response) {
	jData, err := json.Marshal(resp)
	if err != nil {
		BadRequest(w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(jData)
	if err != nil {
		return
	}
}

func OKResponse(w http.ResponseWriter, resp Response) {
	jData, err := json.Marshal(resp)
	if err != nil {
		BadRequest(w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jData)
	if err != nil {
		return
	}
}

func DeleteResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func NotFound(w http.ResponseWriter) {
	resp, err := json.Marshal(Response{Message: "404 Not found"})
	if err != nil {
		BadRequest(w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	_, err = w.Write(resp)
	return
}
