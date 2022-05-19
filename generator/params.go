package generator

import (
	"errors"
	"math"
)

var (
	ErrCountTooMuch    = errors.New("count value too much")
	ErrEmptyOutputFile = errors.New("no output file specified")
)

type Params struct {
	Count   int
	OutFile string
	Format  string
	Verbose bool
}

func (p *Params) Validate() error {
	switch {
	case p.Count > math.MaxInt:
		return ErrCountTooMuch
	case len(p.OutFile) == 0:
		return ErrEmptyOutputFile
	case FileFormat(p.Format) != Csv && FileFormat(p.Format) != Json:
		p.Format = string(Csv)
		return nil
	}

	return nil
}
