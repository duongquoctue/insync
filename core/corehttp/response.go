package corehttp

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type ResponseErr struct {
	Error string `json:"error"`
}

func RespondJSON(w http.ResponseWriter, httpStatusCode int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(httpStatusCode)
	w.Write(data)
}

func RespondErr(w http.ResponseWriter, httpStatusCode int, message string) {
	if httpStatusCode < 400 {
		httpStatusCode = http.StatusInternalServerError
	}

	data := ResponseErr{
		Error: message,
	}

	bytes, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Content-Length", strconv.Itoa(len(bytes)))
	w.WriteHeader(httpStatusCode)
	w.Write(bytes)
}
