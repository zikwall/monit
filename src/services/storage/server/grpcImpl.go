package server

import (
	"context"
	"fmt"

	clickhousebuffer "github.com/zikwall/clickhouse-buffer"
	"github.com/zikwall/clickhouse-buffer/src/buffer/memory"

	"github.com/zikwall/monit/src/pkg/logger"
	"github.com/zikwall/monit/src/protobuf/common"
	"github.com/zikwall/monit/src/protobuf/storage"
)

type GrpcServerImpl struct {
	storage.UnimplementedStorageServer
	heatmapWriter clickhousebuffer.Writer
}

func NewGRPCServerImpl(buffer clickhousebuffer.Client) *GrpcServerImpl {
	s := &GrpcServerImpl{}
	s.initWriterAPI(buffer)

	return s
}

func (s *GrpcServerImpl) initWriterAPI(buffer clickhousebuffer.Client) {
	s.heatmapWriter = buffer.Writer(
		clickhousebuffer.View{
			Name:    getHeatmapTableName(),
			Columns: getHeatmapColumns(),
		},
		memory.NewBuffer(
			buffer.Options().BatchSize(),
		),
	)

	heatmapErrors := s.heatmapWriter.Errors()
	go func() {
		for err := range heatmapErrors {
			logger.Warning(fmt.Sprintf("[HEATMAP] clickhouse write error: %s\n", err.Error()))
		}
	}()
}

func (s *GrpcServerImpl) WriteHeatmap(
	_ context.Context,
	in *storage.HeatmapMessage,
) (
	*common.EmptyResponse,
	error,
) {
	logger.Info("heatmap message receive")
	logger.Info(in)

	s.heatmapWriter.WriteRow(wrapHeatmap(in))
	return &common.EmptyResponse{}, nil
}

func (s *GrpcServerImpl) WriteMetric(
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
