package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
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
			type document struct {
				Data interface{} `json:"data"`
			}
			buf, _ := ioutil.ReadAll(r.Body)
			rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
			rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf))
			decoder := json.NewDecoder(rdr1)
			r.Body = rdr2
			doc := document{}
			decoder.Decode(&doc)
			_, ok := doc.Data.([]interface{})
			if ok {
				p, err = ctrl.BatchAdd(r)
				break
			}
			p, err = ctrl.Add(r)
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
