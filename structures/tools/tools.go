package tools

func Ptr[T comparable](val T) *T {
	tmp := val
	return &tmp
}
