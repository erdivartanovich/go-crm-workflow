package paginator

import (
	"fmt"
	"math"
	"net/url"
)

type Options struct {
	QueryParameter url.Values
	Path           string
}

type LengthAwareOffsetPaginator struct {
	items  interface{}
	total  int
	limit  int
	offset int
	opt    *Options
}

// NewLengthAwareOffsetPaginator NewLengthAwareOffsetPaginator provide length aware offset paginator
func NewLengthAwareOffsetPaginator(items interface{}, total int, limit int, offset int, opt *Options) *LengthAwareOffsetPaginator {

	p := &LengthAwareOffsetPaginator{
		items:  items,
		total:  total,
		limit:  limit,
		offset: offset,
		opt:    opt,
	}
	return p
}

func (p *LengthAwareOffsetPaginator) LastPage() int {
	totalPage := math.Ceil(float64(p.total) / float64(p.limit))
	last := (totalPage - 1) * float64(p.limit)
	return int(last)
}

func (p *LengthAwareOffsetPaginator) URL(offset int) string {
	if p.opt.QueryParameter == nil {
		qp := url.Values{}
		p.opt.QueryParameter = qp
	}

	p.opt.QueryParameter.Set("page[limit]", fmt.Sprint(p.limit))
	p.opt.QueryParameter.Set("page[offset]", fmt.Sprint(offset))

	uri := p.opt.Path

	if uri == "" {
		uri = "/"
	}

	ret, _ := url.QueryUnescape(fmt.Sprintf("%s?%s", uri, p.opt.QueryParameter.Encode()))
	return ret

}

func (p *LengthAwareOffsetPaginator) LastPageURL() string {
	return p.URL(p.LastPage())
}

func (p *LengthAwareOffsetPaginator) NextPageURL() string {
	next := p.offset + p.limit
	if next >= p.total {
		return ""
	}

	return p.URL(next)
}

func (p *LengthAwareOffsetPaginator) PreviousPageURL() string {
	if p.offset == 0 {
		return ""
	}
	prev := p.offset - p.limit
	if prev < 0 {
		return ""
	}
	return p.URL(prev)
}

func (p *LengthAwareOffsetPaginator) Items() interface{} {

	return p.items
}

func (p *LengthAwareOffsetPaginator) Total() int {
	return p.total
}

func (p *LengthAwareOffsetPaginator) Count() int {
	return int(len(p.items.([]interface{})))
}
