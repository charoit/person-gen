package generator

import (
	"errors"
	"github.com/charoit/person-gen/pkg/faker"
	"math"
)

var (
	ErrCountTooMuch    = errors.New("count value too much")
	ErrEmptyOutputFile = errors.New("no output file specified")
)

type Params struct {
	Count   int
	Sex     string
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
	p.Sex = faker.ParseSex(p.Sex).String()

	return nil
}
