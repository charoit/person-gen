package main

import (
	"os"

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
	// ToolVersion see (http://semver.org/)
	// `-ldflags "-X main.Version=${VERSION}"`.
	ToolVersion = "0.0.0-develop"
)

func main() {
	var logVerbose bool
	var logFormat string

	log := logrus.New()

	app := &cli.App{
		Name:        ToolName,
		Usage:       ToolUsage,
		Version:     ToolVersion,
		Description: ToolDescription,

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "count",
				Aliases:  []string{"c"},
				Usage:    "Count of person `VALUE`",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "out",
				Aliases:  []string{"o"},
				Usage:    "Output result to `FILE`",
				Required: true,
			},
			&cli.StringFlag{
				Name:        "format",
				Aliases:     []string{"f"},
				Usage:       "Output file format `csv|json`",
				Value:       "csv",
				Destination: &logFormat,
				//Required:    true,
			},
			&cli.BoolFlag{
				Name:        "log",
				Aliases:     []string{"l"},
				Usage:       "Output verbose log to console",
				Value:       false,
				Destination: &logVerbose,
			},
		},

		Action: func(c *cli.Context) error {
			log.Info("Starting generation...")
			log.Infof("Verbose: %v, format: %v", logVerbose, logFormat)
			//	dmn := daemon.NewDaemon(cfg, log.NewEntry())
			//	defer func() {
			//		if err := dmn.Close(); err != nil {
			//			log.NewEntry().Errorf("close daemon err: %s", err)
			//		}
			//	}()
			//
			//	return dmn.Run()
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

	log.Info("Success.")
}
