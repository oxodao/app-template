package routes

import (
	"github.com/gorilla/mux"
	"github.com/oxodao/app-template/services"
	"net/http"
)

func Main(prv *services.Provider, r *mux.Router) {
	r.HandleFunc("", indexPage(prv))
}

func indexPage(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	}
}
