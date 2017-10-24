package api

import (
	"net/http"
)

type ResourceCtrl interface {
	Browse(r *http.Request) (Responder, error)
	Read(id string, r *http.Request) (Responder, error)
	Replace(id string, r *http.Request) (Responder, error)
	Edit(id string, r *http.Request) (Responder, error)
	Add(r *http.Request) (Responder, error)
	Delete(id string, r *http.Request) (Responder, error)
	BatchAdd(r *http.Request) (Responder, error)
	BatchEdit(r *http.Request) (Responder, error)
	// BatchReplace(r *http.Request) (Responder, error)
	Destroy(r *http.Request) (Responder, error)
}

type BrowseAdapter interface {
}

type Headers map[string]string

type Middleware interface {
	Handle(r *http.Request, next MiddlewareClosure) error
}

type MiddlewareClosure func(r *http.Request) error

type Responder interface {
	Metadata() map[string]interface{}
	StatusCode() int
	Headers() map[string]string
	WriteResponse(w http.ResponseWriter, err error, r *http.Request)
}
