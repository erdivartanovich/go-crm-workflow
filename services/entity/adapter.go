package entity

import (
	"bytes"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/jinzhu/gorm"
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
	if s == "" {
		return
	}
	a.Include = strings.Split(s, ",")
}

func (adapter *SearchAdapter) ApplySearchAdapter(tx *gorm.DB) *gorm.DB {

	tx = adapter.applyInclude(tx)
	tx = adapter.applyFilters(tx)
	tx = adapter.applySorter(tx)
	tx = adapter.applyPager(tx)

	return tx
}

func (adapter *SearchAdapter) applyInclude(tx *gorm.DB) *gorm.DB {

	if len(adapter.Include) > 0 {

		for _, resource := range adapter.Include {

			tx = tx.Preload(strcase.ToCamel(resource))

		}
	}

	return tx
}

func (adapter *SearchAdapter) applyPager(tx *gorm.DB) *gorm.DB {
	limit := 10
	offset := 0
	if adapter.Page != nil {
		limit = adapter.Page.Limit
		offset = adapter.Page.Offset
	}
	tx = tx.Limit(limit).Offset(offset)
	return tx
}

func (adapter *SearchAdapter) applyFilters(tx *gorm.DB) *gorm.DB {
	if adapter.Filters != nil && len(adapter.Filters) > 0 {

	}
	return tx
}

func (adapter *SearchAdapter) applySorter(tx *gorm.DB) *gorm.DB {
	if adapter.Sort != "" {
		tx = tx.Order(string(adapter.Sort))
	}
	return tx
}

func ToSearchAdapter(val url.Values) *SearchAdapter {
	adapter := &SearchAdapter{}
	adapter.FromURLValues(val)
	return adapter
}
