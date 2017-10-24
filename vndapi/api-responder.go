package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/manyminds/api2go/jsonapi"
)

type ApiResponder struct {
	Data     interface{}
	Code     int
	Meta     map[string]interface{}
	Hostname string
}

func (res *ApiResponder) Metadata() map[string]interface{} {
	return res.Meta
}

func (res *ApiResponder) Headers() map[string]string {
	header := make(map[string]string, 1)
	header["content-type"] = "vnd.api+json"
	return header
}

func (res *ApiResponder) StatusCode() int {
	return res.Code
}

func (res *ApiResponder) WriteResponse(w http.ResponseWriter, err error, r *http.Request) {
	headers := res.Headers()
	if len(headers) > 0 {
		for key, header := range headers {
			w.Header().Set(key, header)
		}
	}

	w.WriteHeader(res.StatusCode())
	document, err := jsonapi.MarshalToStruct(res.Data, &ServerInfo{
		baseUrl: r.URL.String(),
		prefix:  r.Host,
	})

	if err != nil {
		w.Write([]byte(err.Error()))
	}

	if document.Links == nil {
		document.Links = jsonapi.Links{}
		document.Links["first"] = jsonapi.Link{
			Href: fmt.Sprintf("%s%s", r.URL.Host, r.URL.String()),
		}
		document.Links["last"] = jsonapi.Link{
			Href: r.URL.String(),
		}
		document.Links["prev"] = jsonapi.Link{
			Href: r.URL.String(),
		}
		document.Links["next"] = jsonapi.Link{
			Href: r.URL.String(),
		}
	}

	data, err := json.Marshal(document)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	w.Write(data)
}

type ServerInfo struct {
	baseUrl string
	prefix  string
}

func (i *ServerInfo) GetBaseURL() string {
	fmt.Println(i.baseUrl)
	return i.baseUrl
}

func (i *ServerInfo) GetPrefix() string {
	return i.prefix
}
