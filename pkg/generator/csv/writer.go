package csv

import (
	"io"

	"github.com/trimmer-io/go-csv"
)

type Writer struct {
	encoder *csv.Encoder
}

func New(w io.Writer) *Writer {
	return &Writer{
		encoder: csv.NewEncoder(w),
	}
}

func (w *Writer) Append(v interface{}) error {
	w.encoder.Separator(';')
	if !w.encoder.HeaderWritten() {
		if err := w.encoder.EncodeHeader(nil, v); err != nil {
			return err
		}
	}

	return w.encoder.EncodeRecord(v)
}
