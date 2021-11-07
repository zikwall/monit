package server

import (
	"github.com/zikwall/clickhouse-buffer/src/buffer"
	"github.com/zikwall/monit/src/protobuf/storage"
)

func getHeatmapTableName() string {
	return "monit.heatmap"
}

func getHeatmapColumns() []string {
	return []string{
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
	}
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
