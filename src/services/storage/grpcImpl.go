package main

import (
	"context"
	clickhousebuffer "github.com/zikwall/clickhouse-buffer"
	"github.com/zikwall/clickhouse-buffer/src/buffer"
	"github.com/zikwall/clickhouse-buffer/src/buffer/memory"
	"github.com/zikwall/monit/src/pkg/logger"
	"github.com/zikwall/monit/src/protobuf/common"
	"github.com/zikwall/monit/src/protobuf/storage"
)

type grpcServerImpl struct {
	storage.UnimplementedStorageServer
	heatmapWriter clickhousebuffer.Writer
}

func NewGRPCServerImpl(buffer clickhousebuffer.Client) *grpcServerImpl {
	s := &grpcServerImpl{}
	s.initWriterAPI(buffer)

	return s
}

func (s *grpcServerImpl) initWriterAPI(buffer clickhousebuffer.Client) {
	s.heatmapWriter = buffer.Writer(
		clickhousebuffer.View{
			Name: "monit.heatmap",
			Columns: []string{
				"stream_id",
				"unique_id",
				"platform",
				"app",
				"version",
				"os",
				"browser",
				"country",
				"region",
				"host",
				"user_agent",
				"insert_ts",
				"insert_date",
			},
		},
		memory.NewBuffer(
			buffer.Options().BatchSize(),
		),
	)
}

type heatmap struct {
	StreamID   string
	UniqueID   string
	Platform   string
	App        string
	Version    string
	OS         string
	Browser    string
	Country    string
	Region     string
	Host       string
	UserAgent  string
	InsertTS   string
	InsertDate string
}

func (h *heatmap) Row() buffer.RowSlice {
	return buffer.RowSlice{
		h.StreamID,
		h.UniqueID,
		h.Platform,
		h.App,
		h.Version,
		h.OS,
		h.Browser,
		h.Country,
		h.Region,
		h.Host,
		h.InsertTS,
		h.InsertDate,
	}
}

func wrapHeatmap(h *storage.HeatmapMessage) *heatmap {
	w := &heatmap{
		StreamID: h.StreamID,
		UniqueID: h.UniqueID,
		Platform: h.Platform,
		App:      h.App,
		Version:  h.Version,
		OS:       h.OS,
		Browser:  h.Browser,
		Country:  h.Country,
		Region:   h.Region,
		Host:     h.Host,
	}
	w.InsertTS = ""
	w.InsertDate = ""
	return w
}

func (s *grpcServerImpl) WriteHeatmap(
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

func (s *grpcServerImpl) WriteMetric(
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
