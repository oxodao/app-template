package main

import (
	"github.com/gorilla/mux"
	"github.com/oxodao/app-template/log"
	"github.com/oxodao/app-template/routes"
	"github.com/oxodao/app-template/services"
	"net/http"
	"time"
)

func RunServer(prv *services.Provider) {
	r := mux.NewRouter()

	routes.Api(prv, r.PathPrefix("/api").Subrouter())
	routes.Main(prv, r.PathPrefix("/").Subrouter())

	srv := &http.Server {
		Addr: prv.Config.Web.ListeningHost,
		Handler: r,
		ReadTimeout: 15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
