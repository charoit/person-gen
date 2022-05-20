package generator

type Writer interface {
	Append(v interface{}) error
}
