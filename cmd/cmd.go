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
	},
	Action: func(c *cli.Context) error {

		data, err := pipe.Read()
		if err != nil && !errors.Is(err, pipe.NoPipe) {
			return err
		}
		if data != nil {
			fmt.Println("Data read from pipe")
		}

		if data == nil && c.String("default-content") != EMPTY_CONTENT {
			data = []byte(c.String("default-content"))
		}

		server := server.New(
			c.Int("port"),
			server.ContentType(c.String("content-type")),
			data,
		)
		return server.Start()
	},
}

func Execute(args []string) error {
	fmt.Println("Starting fork...")
	defer fmt.Println("Fork stopped")
	l := logger.Named("cmd")
	err := app.Run(args)
	if err != nil {
		l.Error("Error running command", zap.Error(err))
		return err
	}
	return nil
}
