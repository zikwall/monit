package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/urfave/cli/v2"
	"github.com/zikwall/monit/src/pkg/logger"
	"github.com/zikwall/monit/src/pkg/signal"
	"github.com/zikwall/monit/src/services/api/actions"
	"github.com/zikwall/monit/src/services/api/middlewares"
	"github.com/zikwall/monit/src/services/api/service"
	"net"
	"os"
	"strings"
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
				Name:     "bind-socket",
				Usage:    "Path to unix socket file for UDS listener",
				Required: false,
				Value:    "/tmp/monit_api.sock",
				EnvVars:  []string{"BIND_SOCKET"},
			},
			&cli.IntFlag{
				Name:     "listener",
				Usage:    "UDS or TCP, default TCP",
				Required: false,
				Value:    ListenerTCP,
				EnvVars:  []string{"LISTENER"},
			},
			&cli.StringFlag{
				Name:     "storage-address",
				Required: true,
				Value:    "0.0.0.0:7000",
				Usage:    "Storage gRPC host",
				EnvVars:  []string{"STORAGE_ADDRESS"},
			},
			&cli.StringFlag{
				Name:     "repository-address",
				Required: true,
				Usage:    "Repository gRPC host",
				EnvVars:  []string{"REPOSITORY_ADDRESS"},
			},
			&cli.StringFlag{
				Name:     "monitoring-address",
				Required: true,
				Usage:    "Monitoring gRPC host",
				EnvVars:  []string{"MONITORING_ADDRESS"},
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
	await, stop := signal.Notifier(func() {
		logger.Info("received a system signal to shutdown API server, start the shutdown process..")
	})

	apiService, err := service.New(ctx.Context, &service.ImplOptions{
		StorageAddress: ctx.String("storage-address"),
	})

	if err != nil {
		return err
	}

	defer func() {
		apiService.Shutdown(func(err error) {
			logger.Warning(err)
		})
		apiService.Stacktrace()
	}()

	go func() {
		controller := actions.NewHTTPController(apiService.StorageClient.Client())
		app := fiber.New(fiber.Config{
			ErrorHandler: middlewares.ErrorHandler,
		})

		api := app.Group("/api")
		api.Post("/data/heatmap", controller.Heatmap)
		api.Post("/data/event", controller.Event)

		ln, err := resolveListener(
			apiService.Context,
			ctx.Int("listener"),
			ctx.String("bind-socket"),
			ctx.String("bind-address"),
		)

		if err != nil {
			stop(err)
			return
		}

		if err := app.Listener(ln); err != nil {
			stop(err)
		}
	}()

	logger.Info("run api service")
	return await()
}

func resolveListener(ctx context.Context, listener int, uds, tcp string) (net.Listener, error) {
	if listener == ListenerUDS {
		defer maybeChmodSocket(ctx, uds)
		ln, err := listenToUnix(uds)

		return ln, err
	}

	if !strings.Contains(tcp, ":") {
		tcp = ":" + tcp
	}

	ln, err := net.Listen(fiber.NetworkTCP4, tcp)

	if err != nil {
		return nil, err
	}

	return ln, nil
}
