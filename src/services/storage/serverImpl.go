package main

import (
	"context"
	"github.com/zikwall/monit/src/pkg/logger"
	"github.com/zikwall/monit/src/protobuf/common"
	"github.com/zikwall/monit/src/protobuf/storage"
)

type serverImpl struct {
	storage.UnimplementedStorageServer
}

func (s *serverImpl) WriteHeatmap(
	_ context.Context,
	in *storage.HeatmapMessage,
) (
	*common.EmptyResponse,
	error,
) {
	logger.Info("heatmap message receive")
	logger.Info(in)
	return &common.EmptyResponse{}, nil
}

func (s *serverImpl) WriteMetric(
	_ context.Context,
	in *storage.MetricMessage,
) (
	*common.EmptyResponse,
	error,
) {
	logger.Info("metric message receive")
	logger.Info(in)
	return &common.EmptyResponse{}, nil
}
