package paginator

type Paginator interface {
	LastPage() int
	URL(offset int) string
	LastPageURL() string
	NextPageURL() string
	PreviousPageURL() string
	Items() interface{}
	Total() int
	Count() int
}
