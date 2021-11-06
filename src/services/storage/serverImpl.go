package main

import (
	"context"
	"github.com/zikwall/monit/src/pkg/logger"
	"github.com/zikwall/monit/src/protobuf/storage"
)

type serverImpl struct {
	storage.UnimplementedStorageServer
}

func (s *serverImpl) WriteHeatmap(
	_ context.Context,
	in *storage.HeatmapMessage,
) (
	*storage.Response,
	error,
) {
	logger.Info("heatmap message receive")
	logger.Info(in)
	return &storage.Response{}, nil
}

func (s *serverImpl) WriteMetric(
	_ context.Context,
	in *storage.MetricMessage,
) (
	*storage.Response,
	error,
) {
	logger.Info("metric message receive")
	logger.Info(in)
	return &storage.Response{}, nil
}
