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
				Usage:    "Run service in host",
				EnvVars:  []string{"BIND_ADDRESS"},
			},
			&cli.StringFlag{
				Name:     "database-host",
				Required: true,
				Usage:    "Database host",
				EnvVars:  []string{"DATABASE_HOST"},
				FilePath: "/srv/monit_secret/database_host",
			},
			&cli.StringFlag{
				Name:     "database-user",
				Required: true,
				Usage:    "Database user",
				EnvVars:  []string{"DATABASE_USER"},
				FilePath: "/srv/monit_secret/database_user",
			},
			&cli.StringFlag{
				Name:     "database-password",
				Required: true,
				Usage:    "Database password",
				EnvVars:  []string{"DATABASE_PASSWORD"},
				FilePath: "/srv/monit_secret/database_password",
			},
			&cli.StringFlag{
				Name:     "database-name",
				Required: true,
				Usage:    "Database name",
				EnvVars:  []string{"DATABASE_NAME"},
				FilePath: "/srv/monit_secret/database_name",
			},
			&cli.StringFlag{
				Name:     "database-dialect",
				Required: true,
				Usage:    "Database dialect: mysql, postgres, sqlite3, sqlserver etc",
				EnvVars:  []string{"DATABASE_DIALECT"},
				FilePath: "/srv/monit_secret/database_dialect",
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
		logger.Info("received a system signal to shutdown REPOSITORY server, start the shutdown process..")
	})

	go func() {
		// run gRPC server here
	}()

	return await()
}
