package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

func handleResourceRoute(path string, ctrl ResourceCtrl, middlewares ...Middleware) {

}

func handleBulkResourceServiceRoute(ctrl ResourceCtrl, middlewares ...Middleware) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			p   Responder
			err error
		)
		switch r.Method {
		case http.MethodGet:
			p, err = ctrl.Browse(r)
			break
		case http.MethodPost:
			p, err = ctrl.BatchAdd(r)
			break
		case http.MethodPatch:
			p, err = ctrl.BatchEdit(r)
			break
		}
		p.WriteResponse(w, err, r)
	}
}

func handleResourceServiceRoute(ctrl ResourceCtrl, middlewares ...Middleware) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var (
			p   Responder
			err error
		)
		vars := mux.Vars(r)
		id := vars["id"]

		switch r.Method {
		case http.MethodGet:
			p, err = ctrl.Read(id, r)
			break
		case http.MethodPost:
			p, err = ctrl.Replace(id, r)
			break
		case http.MethodPatch:
			p, err = ctrl.Edit(id, r)
			break
		}

		p.WriteResponse(w, err, r)
	}
}
