package paginator

type Paginator interface {
	LastPage() uint
	URL(offset uint) string
	LastPageURL() string
	NextPageURL() string
	PreviousPageURL() string
	Items() []interface{}
	Total() uint
	Count() uint
}
