package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	paginator "github.com/kwri/go-workflow/gorm-paginator"
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

	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte(err.Error()))
		return
	}
	headers := res.Headers()
	if len(headers) > 0 {
		for key, header := range headers {
			w.Header().Set(key, header)
		}
	}

	w.WriteHeader(res.StatusCode())

	items := res.Data
	paginator, isPaginator := items.(paginator.Paginator)
	if isPaginator {
		items = paginator.Items()
	}

	document, err := jsonapi.MarshalToStruct(items, &ServerInfo{
		baseUrl: r.URL.String(),
		prefix:  r.Host,
	})

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	ensureDocumentLinks(document, paginator, isPaginator)
	document.Meta = res.Metadata()
	data, err := json.Marshal(document)

	if err != nil {
		w.Write([]byte(err.Error()))

		return
	}

	w.Write(data)
}

func ensureDocumentLinks(document *jsonapi.Document, paginator paginator.Paginator, isPaginator bool) {

	if document.Links == nil {
		document.Links = jsonapi.Links{}
	}

	if isPaginator {
		document.Links["first"] = jsonapi.Link{
			Href: paginator.URL(0),
		}
		document.Links["last"] = jsonapi.Link{
			Href: paginator.LastPageURL(),
		}
		document.Links["prev"] = jsonapi.Link{
			Href: paginator.PreviousPageURL(),
		}
		document.Links["next"] = jsonapi.Link{
			Href: paginator.NextPageURL(),
		}
	}
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
