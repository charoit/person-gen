package main

import (
	"embed"
	"fmt"
	"github.com/charoit/person-gen/pkg/faker"

	"os"
	"time"

	"github.com/charoit/person-gen/pkg/generator"
	"github.com/charoit/person-gen/pkg/storage/files"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

const (
	// ToolName name of tool.
	ToolName = "person-gen"

	// ToolUsage short tool description.
	ToolUsage = "person generator"

	// ToolDescription full tool description.
	ToolDescription = "Person generator for test data."
)

var (
	// Version see (http://semver.org/)
	// -ldflags "-X main.Version=${VERSION}"
	Version = "0.0.0-develop"
)

//go:embed resources/*
var resources embed.FS

func main() {
	var fakeData faker.Fake
	var elapsedTime time.Time

	log := &logrus.Logger{
		Out:   os.Stderr,
		Level: logrus.InfoLevel,
		Formatter: &logrus.TextFormatter{
			DisableTimestamp: true,
		},
	}

	params := generator.Params{}
	app := &cli.App{
		Name:        ToolName,
		Usage:       ToolUsage,
		Version:     Version,
		Description: ToolDescription,

		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "count",
				Aliases:     []string{"c"},
				Usage:       "Count of person `VALUE`",
				Required:    true,
				Destination: &params.Count,
			},
			&cli.StringFlag{
				Name:        "sex",
				Aliases:     []string{"s"},
				Usage:       "Sex of person `VALUE`",
				Value:       "all",
				Destination: &params.Sex,
			},
			&cli.StringFlag{
				Name:        "out",
				Aliases:     []string{"o"},
				Usage:       "Output result to `FILE`",
				Required:    true,
				Destination: &params.OutFile,
			},
			&cli.StringFlag{
				Name:        "format",
				Aliases:     []string{"f"},
				Usage:       "Output file format `csv|json`",
				Value:       "csv",
				Destination: &params.Format,
			},
			&cli.BoolFlag{
				Name:        "log",
				Aliases:     []string{"l"},
				Usage:       "Output verbose log to console",
				Value:       false,
				Destination: &params.Verbose,
			},
		},

		Before: func(c *cli.Context) error {
			var err error
			if fakeData, err = files.NewStorage(resources).Load(); err != nil {
				return fmt.Errorf("loading data failed: %w", err)
			}

			return nil
		},

		Action: func(c *cli.Context) error {
			if err := params.Validate(); err != nil {
				return fmt.Errorf("validate failed: %w", err)
			}

			log.Info("Starting generation...")
			elapsedTime = time.Now()

			gen := generator.New(log, &params)
			total, err := gen.Generate(&fakeData)
			if err != nil {
				return fmt.Errorf("generated failed: %w", err)
			}

			log.Info("-------------------------------------------")
			log.Info("Total person generated: ", total)
			log.Info("Total time elapsed: ", time.Since(elapsedTime))
			log.Info("Generated success.")

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
