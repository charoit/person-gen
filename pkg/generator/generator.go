package generator

import (
	"io"
	"math"
	"os"
	"strconv"

	"github.com/charoit/person-gen/pkg/faker"
	"github.com/charoit/person-gen/pkg/generator/csv"
	"github.com/charoit/person-gen/pkg/generator/json"
	"github.com/sirupsen/logrus"
)

type FileFormat string

const (
	Csv  FileFormat = "csv"
	Json FileFormat = "json"
)

type Generator struct {
	log           *logrus.Logger
	params        *Params
	verboseFormat string
}

func New(log *logrus.Logger, params *Params) *Generator {
	n := int(math.Ceil(math.Log10(float64(params.Count) + 0.5)))

	return &Generator{
		log:           log,
		params:        params,
		verboseFormat: "%" + strconv.Itoa(n) + "d: %v",
	}
}

func (g *Generator) Generate(data *faker.Fake) (int, error) {
	file, err := os.OpenFile(g.params.OutFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return 0, err
	}

	defer func() {
		if err = file.Close(); err != nil {
			g.log.Error(err)
		}
	}()

	sex := faker.ParseSex(g.params.Sex)
	wr := g.writer(file, FileFormat(g.params.Format))
	gn := faker.New(data)

	total := 0
	for total != g.params.Count {
		total++
		person := gn.MakePerson("rus", sex)
		if err = wr.Append(person); err != nil {
			return total, err
		}

		if g.params.Verbose {
			g.log.Infof(g.verboseFormat, total, person)
		}
	}

	return total, nil
}

func (g *Generator) writer(w io.Writer, format FileFormat) Writer {
	var wr Writer
	switch format {
	case Csv:
		wr = csv.New(w)
	case Json:
		wr = json.New(w)
	}

	return wr
}
