package filter

type PhotoFilter interface {
	Filter(int, int) bool
}
