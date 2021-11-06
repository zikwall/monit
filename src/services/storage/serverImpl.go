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
	*storage.HeatmapResponse,
	error,
) {
	logger.Info(in)
	return &storage.HeatmapResponse{Error: &common.Error{Code: 0, Message: ""}}, nil
}
