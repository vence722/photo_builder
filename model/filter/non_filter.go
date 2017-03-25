package filter

type NonFilter struct{}

func NewNonFilter() *NonFilter {
	return &NonFilter{}
}

func (this *NonFilter) Filter(x int, y int) bool {
	return true
}
