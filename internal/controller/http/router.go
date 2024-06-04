package http

import (
	"fmt"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"net/http"
)

func NewRouter(c Controller, port int) http.Handler {
	r := mux.NewRouter()

	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", port)), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	r.HandleFunc("/currency/save/{date}", func(w http.ResponseWriter, r *http.Request) {
		//TODO set content-type in middleware
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		dateString := params["date"]
		c.Save(r.Context(), w, r, dateString)
	}).Methods("POST")

	r.HandleFunc("/currency/{date}/{code}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		dateString := params["date"]
		code := params["code"]
		c.List(r.Context(), w, r, dateString, code)
	}).Methods("GET")

	return r
}
