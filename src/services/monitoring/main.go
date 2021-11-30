package main

import (
	"os"

	"github.com/urfave/cli/v2"

	"github.com/zikwall/monit/src/pkg/logger"
	"github.com/zikwall/monit/src/pkg/signal"
)

func main() {
	application := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "bind-address",
				Required: true,
				Usage:    "Run service in hosts",
				EnvVars:  []string{"BIND_ADDRESS"},
			},
			&cli.StringFlag{
				Name:     "storage-address",
				Required: true,
				Usage:    "Storage gRPC host",
				EnvVars:  []string{"STORAGE_ADDRESS"},
			},
			&cli.StringFlag{
				Name:     "repository-address",
				Required: true,
				Usage:    "Repository gRPC hosts",
				EnvVars:  []string{"REPOSITORY_ADDRESS"},
			},
			&cli.StringFlag{
				Name:     "api-address",
				Required: true,
				Usage:    "Monitoring gRPC host",
				EnvVars:  []string{"API_ADDRESS"},
			},
			&cli.BoolFlag{
				Name:    "debug",
				EnvVars: []string{"DEBUG"},
				Value:   false,
			},
		},
		Action: Main,
	}

	if err := application.Run(os.Args); err != nil {
		logger.Error(err)
	}
}

func Main(_ *cli.Context) error {
	await, _ := signal.Notifier(func() {
		logger.Info("received a system signal to shutdown API server, start the shutdown process..")
	})

	go func() {
		// run gRPC server here
	}()

	return await()
}
