package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/castillobg/pong/core"
	"github.com/gorilla/mux"
)

func Listen(port int) error {
	r := mux.NewRouter()
	r.Path("/api/pongs").Methods(http.MethodGet).HandlerFunc(pongsHandler)
	return http.ListenAndServe(":"+strconv.Itoa(port), r)
}

func pongsHandler(w http.ResponseWriter, _ *http.Request) {
	response := struct{ Pongs int }{Pongs: core.Pongs()}
	jsonBytes, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Write(jsonBytes)
}
