package json

import (
	"encoding/json"
	"io"
)

type Writer struct {
	encoder *json.Encoder
}

func New(w io.Writer) *Writer {
	return &Writer{
		encoder: json.NewEncoder(w),
	}
}

func (w *Writer) Append(v interface{}) error {
	return w.encoder.Encode(v)
}
