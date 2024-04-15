package cmd

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"pretender/internal/server"
	"pretender/pkg/logger"
	"pretender/pkg/pipe"
)

const (
	EMPTY_CONTENT = ""
)

var app = &cli.App{
	Name:  "pretender",
	Usage: "A tool to generate fake server responses",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "port",
			Aliases: []string{"p"},
			Usage:   "The port to listen on",
			Value:   "8080",
		},
		&cli.StringFlag{
			Name:    "content-type",
			Aliases: []string{"c"},
			Usage:   "The content type to use",
			Value:   server.JSONContentType.String(),
		},
		&cli.StringFlag{
			Name:    "default-content",
			Aliases: []string{"d"},
			Usage:   "The default content to return",
			Value:   EMPTY_CONTENT,
		},
		&cli.BoolFlag{
			Name:    "prefork",
			Aliases: []string{"f"},
			Usage:   "Prefork the server",
			Value:   false,
		},
	},
	Action: func(c *cli.Context) error {
		if !c.Bool("prefork") {
			fmt.Println("Starting fork...")
			defer fmt.Println("Fork stopped")
		} else {
			fmt.Println("Preforking...")
			defer fmt.Println("Prefork stopped")
		}
		data, err := pipe.Read()
		if err != nil && !errors.Is(err, pipe.NoPipe) {
			return err
		}

		if c.Bool("prefork") && data != nil {
			fmt.Println("Reading from pipe does not work with prefork")
			return nil
		}

		if data == nil && c.String("default-content") != EMPTY_CONTENT {
			data = []byte(c.String("default-content"))
		}

		server := server.New(
			c.Int("port"),
			server.ContentType(c.String("content-type")),
			data,
			c.Bool("prefork"),
		)
		return server.Start()
	},
}

func Execute(args []string) error {

	l := logger.Named("cmd")
	err := app.Run(args)
	if err != nil {
		l.Error("Error running command", zap.Error(err))
		return err
	}
	return nil
}
