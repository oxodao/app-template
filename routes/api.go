package routes

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/oxodao/app-template/middlewares"
	"github.com/oxodao/app-template/services"
	"net/http"
)

func Api(prv *services.Provider, r *mux.Router) {
	r.Use(middlewares.Json)

	r.HandleFunc("/hello", apiHelloPage(prv))
}

func apiHelloPage(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		str, err := json.Marshal(struct {
			Hello string
		} {
			Hello: "world",
		})

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		w.Write(str)
	}
}
