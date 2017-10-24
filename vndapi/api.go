package api

import (
	"fmt"
	"net/http"
)

type Api struct {
	versionPath string
}

func (a *Api) Resource(name string, ctrl ResourceCtrl, middlewares ...Middleware) {
	path := fmt.Sprintf("/%s/%s", a.versionPath, name)
	router.HandleFunc(
		path,
		handleBulkResourceServiceRoute(path, ctrl, middlewares...),
	).Methods(http.MethodGet, http.MethodPost, http.MethodPatch)
	router.HandleFunc(
		fmt.Sprintf("%s/:id", path),
		handleResourceServiceRoute(path, ctrl, middlewares...),
	).Methods(http.MethodGet, http.MethodPost, http.MethodPatch)

}

func New(versionPath string) *Api {
	return &Api{
		versionPath: versionPath,
	}
}

func ListenAndServe(addr string) {
	http.ListenAndServe(addr, router)
}

func ListenAndServeTLS(addr string, certFile string, keyFile string) {
	http.ListenAndServeTLS(addr, certFile, keyFile, nil)
}

func validateMiddlewares(index int, r *http.Request, mdws ...Middleware) error {

	if count := len(mdws); count <= 0 || index >= count {
		return nil
	}

	next := func(req *http.Request) error {
		return validateMiddlewares(index+1, req, mdws...)
	}

	return mdws[index].Handle(r, next)
}
