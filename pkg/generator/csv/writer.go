package csv

import (
	"fmt"
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
		if _, err := w.encoder.Write([]byte{0xEF, 0xBB, 0xBF}); err != nil {
			return fmt.Errorf("write BOM: %w", err)
		}

		if err := w.encoder.EncodeHeader(nil, v); err != nil {
			return fmt.Errorf("encode CSV header: %w", err)
		}
	}

	return w.encoder.EncodeRecord(v)
}
