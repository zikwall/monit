package storage

import (
	"github.com/urfave/cli/v2"
	"github.com/zikwall/monit/src/pkg/logger"
	"github.com/zikwall/monit/src/pkg/signal"
	"os"
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
				Name:     "clickhouse-address",
				Usage:    "Clickhouse server address",
				Required: true,
				EnvVars:  []string{"CLICKHOUSE_ADDRESS"},
				FilePath: "/srv/monit_secret/clickhouse_address",
			},
			&cli.StringFlag{
				Name:     "clickhouse-user",
				Usage:    "Clickhouse server user",
				EnvVars:  []string{"CLICKHOUSE_USER"},
				FilePath: "/srv/monit_secret/clickhouse_user",
			},
			&cli.StringFlag{
				Name:     "clickhouse-password",
				Usage:    "Clickhouse server user password",
				EnvVars:  []string{"CLICKHOUSE_PASSWORD"},
				FilePath: "/srv/monit_secret/clickhouse_password",
			},
			&cli.StringFlag{
				Name:     "clickhouse-database",
				Usage:    "Clickhouse server database name",
				EnvVars:  []string{"CLICKHOUSE_DATABASE"},
				FilePath: "/srv/monit_secret/clickhouse_database",
			},
			&cli.StringFlag{
				Name:     "clickhouse-alt-hosts",
				Usage:    "Comma separated list of single address host for load-balancing",
				EnvVars:  []string{"CLICKHOUSE_ALT_HOSTS"},
				FilePath: "/srv/monit_secret/clickhouse_alt_hosts",
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

func Main(ctx *cli.Context) error {
	await, _ := signal.Notifier(func() {
		logger.Info("received a system signal to shutdown STORAGE server, start the shutdown process..")
	})

	go func() {
		// run gRPC server here
	}()

	return await()
}
