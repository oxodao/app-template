package routes

import (
	"github.com/gorilla/mux"
	"github.com/oxodao/app-template/cerrors"
	"github.com/oxodao/app-template/services"
	"net/http"
)

func Main(prv *services.Provider, r *mux.Router) {
	r.HandleFunc("", indexPage(prv))
	r.HandleFunc("/{user}", helloUserPage(prv))
}

func indexPage(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	}
}

func helloUserPage(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		if len(vars["user"]) == 0 {
			w.WriteHeader(http.StatusBadRequest)
		}

		user, err := prv.Dal.User.FindByUsername(vars["user"])
		if cerrors.WriteError(w, err) {
			return
		}

		w.Write([]byte("Hello " + user.Username))
	}
}
