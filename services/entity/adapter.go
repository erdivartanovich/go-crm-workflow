package entity

import (
	"bytes"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

type Page struct {
	Limit  int
	Offset int
}

type Filter struct {
	Name      string
	Condition string
	Value     string
}

type Sort string

type SearchAdapter struct {
	Include []string
	Page    *Page
	Filters []*Filter
	Sort    Sort
}

func (a *SearchAdapter) FromURLValues(value url.Values) {
	a.UnmarshalFilter(value)

	limit, e := strconv.Atoi(value.Get("page[limit]"))
	if e != nil {
		limit = 10
	}

	offset, e := strconv.Atoi(value.Get("page[offset]"))
	if e != nil {
		offset = 0
	}

	a.Page = &Page{
		Limit:  limit,
		Offset: offset,
	}

	a.UnmarshalSort(value.Get("sort"))
	a.UnmarshalInclude(value.Get("include"))

}

func (a *SearchAdapter) UnmarshalFilter(value url.Values) {
	var Filters []*Filter
	for name, value := range value {
		if strings.HasPrefix(name, "filter") {
			r, _ := regexp.Compile(`\[([a-zA-Z.\-\:\(\)\,]+)\]`)
			a := r.FindAllStringSubmatch(name, 2)
			if len(a) == 2 {
				filter := &Filter{
					Name:      a[0][1],
					Condition: a[1][1],
					Value:     value[len(value)-1],
				}
				Filters = append(Filters, filter)
			}
		}
	}
	a.Filters = Filters

}

func (a *SearchAdapter) UnmarshalSort(s string) {

	sorts := strings.Split(s, ",")
	var buff bytes.Buffer

	if len(sorts) > 0 {
		for _, attrName := range sorts {
			if attrName != "" {
				if len(buff.Bytes()) > 0 {
					buff.WriteString(",")
				}

				if attrName[0:1] == "-" {
					attrName = fmt.Sprintf("%s %s", attrName[1:], "desc")
				}

				buff.WriteString(attrName)
			}
		}
	}

	a.Sort = Sort(buff.String())
}

func (a *SearchAdapter) UnmarshalInclude(s string) {

	a.Include = strings.Split(s, ",")
}
